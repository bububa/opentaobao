package legalsuit

// CourtPartyModel 结构体
type CourtPartyModel struct {
	// 送达情况
	DeliverySituation string `json:"delivery_situation,omitempty" xml:"delivery_situation,omitempty"`
	// 是否应诉
	IsRespond string `json:"is_respond,omitempty" xml:"is_respond,omitempty"`
	// 是否到庭
	IsAppear string `json:"is_appear,omitempty" xml:"is_appear,omitempty"`
	// 身份
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 身份识别
	ProxyIdentity string `json:"proxy_identity,omitempty" xml:"proxy_identity,omitempty"`
	// 代理人名称
	ProxyName string `json:"proxy_name,omitempty" xml:"proxy_name,omitempty"`
	// 是否有代理人
	IsProxy string `json:"is_proxy,omitempty" xml:"is_proxy,omitempty"`
	// 原告/被告名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
