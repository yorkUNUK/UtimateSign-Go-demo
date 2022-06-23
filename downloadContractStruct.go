package main

/*
* @Author: York
* @Date:   2022/6/23 12:01 PM
 */

type DownloadContract struct {
	Operator         Operator `json:"operator,omitempty"`
	ContractIds      []string `json:"contractIds"`
	FileType         string   `json:"fileType,omitempty"`
	EncodeByBase64   bool     `json:"encodeByBase64"`
	SignerAttachment bool     `json:"signerAttachment"`
}

type Operator struct {
	EnterpriseName string `json:"enterpriseName,omitempty"`
	Account        string `json:"account,omitempty"`
	BizName        string `json:"bizName,omitempty"`
}
