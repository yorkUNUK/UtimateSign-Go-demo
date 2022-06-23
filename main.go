package main

/*
* @Author: York
* @Date:   2022/6/22 10:16 AM
 */
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

// 请填充！
const (
	PrivateKey   = ""
	ClientId     = ""
	ClientSecret = ""
	Host         = "https://api.bestsign.info"
)

func main() {
	/*************************** 发送合同 ***************************/
	sendTemplateContract := new(SendTemplateContract)
	//sendTemplateContract.TemplateId = "2902275519070340097"
	/* ROLES */
	roles := []Role{
		{
			RoleId: "3062952243905954823",
			UserInfo: UserInfo{
				UserAccount:    "15867397177",
				EnterpriseName: "xyc公司",
			},
		},
		{
			UserInfo: UserInfo{
				UserAccount:    "18236922636",
				EnterpriseName: "上上签签签公司",
			},
			ReceiverType: "SIGNER",
			UserType:     "ENTERPRISE",
			RoleName:     "ysp",
		},
	}
	sendTemplateContract.Roles = roles
	/* DOCUMENTS */

	documents := []Document{
		{
			DocumentId: "2902276315971322882",
			ContractConfig: ContractConfig{
				ContractTitle: "Go测试1",
			},
			AppendingSignLabels: []AppendingSignLabels{
				{
					X:          0.8,
					Y:          0.8,
					RoleName:   "ysp",
					Type:       "SEAL",
					PageNumber: 1,
				},
			},
		},
		{
			Content:  readFileAndConvert2Base64("/Users/edianyun/Downloads/york.pdf"),
			FileName: "Go_test.pdf",
			ContractConfig: ContractConfig{
				ContractTitle: "Go测试2",
			},
			AppendingSignLabels: []AppendingSignLabels{
				{
					X:          0.2,
					Y:          0.2,
					RoleName:   "公司",
					Type:       "SEAL",
					PageNumber: 1,
				},
				{
					X:          0.8,
					Y:          0.8,
					RoleName:   "ysp",
					Type:       "SEAL",
					PageNumber: 1,
				},
			},
		},
	}
	sendTemplateContract.Documents = documents
	/*TEXTLABELS */
	textLabels := []TextLabel{
		{Name: "Go测试111", Value: "111"},
		{Name: "Go测试222", Value: "222"},
	}
	sendTemplateContract.TextLabels = textLabels
	sendContractRequest, err := json.Marshal(sendTemplateContract)
	if err != nil {
		fmt.Println("转换为json错误")
		os.Exit(-1)
	}
	sendContractResult := executeRequest("/api/templates/send-contracts-sync-v2", "POST", string(sendContractRequest))

	sendContractResponse := Response{}
	sendContractResponseData := Data{}
	json.Unmarshal([]byte(sendContractResult), &sendContractResponse)
	json.Unmarshal(sendContractResponse.Data, &sendContractResponseData)
	fmt.Println("*********************发送合同*********************")
	fmt.Println(sendContractResult)
	fmt.Println("合同ID为： " + sendContractResponseData.ContractId)
	fmt.Println("*************************************************")
	/***************************************************************/

	/*************************** 下载合同 ***************************/
	downloadContract := DownloadContract{
		ContractIds:    []string{"3080301485075293999"},
		EncodeByBase64: false,
		FileType:       "pdf",
	}
	downloadContractRequest, err := json.Marshal(downloadContract)
	if err != nil {
		fmt.Println("转换为json错误")
		os.Exit(-1)
	}
	downloadContractResult := executeRequest("/api/contracts/download-file", "POST", string(downloadContractRequest))
	file, err := base64.StdEncoding.DecodeString(downloadContractResult)
	fmt.Println("*********************发送合同*********************")
	if err == nil {
		fmt.Println("Download successfully!")
		err2 := ioutil.WriteFile("/Users/edianyun/Downloads/6-23.pdf", file, 0777)
		if err2 != nil {
			log.Fatal(err.Error())
		}
	} else {
		fmt.Println(downloadContractResult)
	}
	fmt.Println("*************************************************")
	/***************************************************************/

	/*************************** 查询模板列表 ***************************/
	uriWithParam := fmt.Sprintf(
		"/api/templates/v2?currentPage=%d&pageSize=%d&account=%s&enterpriseName=%s",
		1,
		20,
		"15867397177",
		url.QueryEscape("徐宇超市场经营管理有限公司"))
	queryTemplatesResult := executeRequest(uriWithParam, "GET", "")
	fmt.Println("*********************查询模板列表*********************")
	fmt.Println(queryTemplatesResult)
	fmt.Println("****************************************************")
	/***************************************************************/
}

func readFileAndConvert2Base64(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}
	result := base64.StdEncoding.EncodeToString(f)
	return result
}
