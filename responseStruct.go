package main

/*
* @Author: York
* @Date:   2022/6/23 12:00 PM
 */
import "encoding/json"

/* 以下response的结构体有待补充 */

type Response struct {
	Code    string          `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

type Data struct {
	ExpiresIn    string      `json:"expiresIn"`
	Expiration   interface{} `json:"expiration"`
	TokenType    string      `json:"tokenType"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	ContractId   string      `json:"contractId"`
}
