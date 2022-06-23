package main

/*
* @Author: York
* @Date:   2022/6/22 10:16 AM
 */
import (
	"bytes"
	"crypto"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var tokenCache string //避免每次调用接口都申请token

func executeRequest(uriWithParam, method, requestBody string) string {
	retryTime := 3
	for retryTime > 0 {
		client := &http.Client{}
		timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		signRSA := signRequest(uriWithParam, timestamp, requestBody)
		token := queryToken()
		req, err := http.NewRequest(
			method,
			fmt.Sprintf("%s%s", Host, uriWithParam),
			strings.NewReader(requestBody))
		if err != nil {
			fmt.Println("创建http-request失败")
			os.Exit(-1)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
		req.Header.Set("bestsign-sign-timestamp", timestamp)
		req.Header.Set("bestsign-client-id", ClientId)
		req.Header.Set("bestsign-signature-type", "RSA256")
		req.Header.Set("bestsign-signature", signRSA)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("http请求失败")
			log.Fatal(err)
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("接口IO失败")
			log.Fatal(err)
		}
		respCode := resp.StatusCode
		if respCode == 401 {
			fmt.Println("token过期")
			invalidToken(tokenCache)
		} else if respCode == 200 {
			if resp.Header.Get("Content-Type") == "application/zip" || resp.Header.Get("Content-Type") == "application/pdf" {
				return base64.StdEncoding.EncodeToString(respBody)
			}
			return string(respBody)
		} else {
			fmt.Printf("Response code is %d", respCode)
			os.Exit(-1)
		}
		retryTime--
	}
	return ""
}

func signRequest(uriWithParam, timestamp, requestBody string) string {
	body := fmt.Sprintf("%x", md5.Sum([]byte(requestBody))) // md5.Sum(data)返回byte类型，fmt.Sprintf()格式化
	content := fmt.Sprintf(
		"bestsign-client-id=%sbestsign-sign-timestamp=%sbestsign-signature-type=%srequest-body=%suri=%s",
		ClientId,
		timestamp,
		"RSA256",
		body,
		uriWithParam)
	sign := RsaSign(content, PrivateKey, crypto.SHA256)
	return sign
}

func queryToken() string {
	if tokenCache == "" {
		client := &http.Client{}
		developInfo := DeveloperInfo{
			ClientId:     ClientId,
			ClientSecret: ClientSecret,
		}
		reqBody, errJson := json.Marshal(developInfo)
		if errJson != nil {
			fmt.Println("转换为json错误")
			os.Exit(-1)
		}
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/api/oa2/client-credentials/token", Host),
			bytes.NewReader(reqBody))
		if err != nil {
			fmt.Println("获取token失败")
			os.Exit(-1)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("获取token时IO失败")
			os.Exit(-1)
		}
		response := Response{}
		data := Data{}
		json.Unmarshal(body, &response)
		json.Unmarshal(response.Data, &data)
		tokenCache = data.AccessToken
		return tokenCache
	} else {
		return tokenCache
	}
	return ""
}

func invalidToken(oldToken string) {
	if oldToken == tokenCache {
		tokenCache = ""
	}
}

type DeveloperInfo struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
