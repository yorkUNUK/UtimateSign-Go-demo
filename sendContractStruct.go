package main

/*
* @Author: York
* @Date:   2022/6/22 10:33 AM
 */

/* 以下request的结构体有待补充 */

type SendTemplateContract struct {
	TemplateId   string      `json:"templateId"`
	Sender       Sender      `json:"sender"`
	Roles        []Role      `json:"roles"`
	Documents    []Document  `json:"documents"`
	TextLabels   []TextLabel `json:"textLabels"`
	ContractName string      `json:"contractName,omitempty"`
	SignOrderly  bool        `json:"signOrderly,omitempty"`
	BizNo        string      `json:"bizNo,omitempty"`
	PushUrl      string      `json:"pushUrl,omitempty"`
}

type Sender struct {
	EnterpriseName string `json:"enterpriseName,omitempty"`
	Account        string `json:"account,omitempty"`
	BizName        string `json:"bizName,omitempty"`
}
type Role struct {
	RoleId       string   `json:"roleId,omitempty"`
	RoleName     string   `json:"roleName,omitempty"`
	Disabled     bool     `json:"disabled,omitempty"`
	ReceiverType string   `json:"receiverType,omitempty"`
	UserType     string   `json:"userType,omitempty"`
	UserInfo     UserInfo `json:"userInfo"`
}

type UserInfo struct {
	EnterpriseName string `json:"enterpriseName,omitempty"`
	UserAccount    string `json:"userAccount,omitempty"`
	UserName       string `json:"userName,omitempty"`
}

type Document struct {
	DocumentId          string                `json:"documentId,omitempty"`
	Disabled            bool                  `json:"disabled,omitempty"`
	Content             string                `json:"content,omitempty"`
	FileName            string                `json:"fileName,omitempty"`
	ContractConfig      ContractConfig        `json:"contractConfig,omitempty"`
	AppendingSignLabels []AppendingSignLabels `json:"appendingSignLabels,omitempty"`
}

type ContractConfig struct {
	ContractType     string `json:"contractType,omitempty"`
	ContractTitle    string `json:"contractTitle,omitempty"`
	CustomContractId string `json:"customContractId,omitempty"`
	SignExpireDays   int    `json:"signExpireDays,omitempty"`
	ContractLifeEnd  int64  `json:"contractLifeEnd,omitempty"`
}

type AppendingSignLabels struct {
	X          float32 `json:"x,omitempty"`
	Y          float32 `json:"y,omitempty"`
	RoleName   string  `json:"roleName,omitempty"`
	Type       string  `json:"type,omitempty"`
	PageNumber int     `json:"pageNumber,omitempty"`
}

type TextLabel struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
