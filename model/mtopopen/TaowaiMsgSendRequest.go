package mtopopen

import (
	"sync"
)

// TaowaiMsgSendRequest 结构体
type TaowaiMsgSendRequest struct {
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 扩展参数
	ExtParams string `json:"ext_params,omitempty" xml:"ext_params,omitempty"`
	// openid
	Openid string `json:"openid,omitempty" xml:"openid,omitempty"`
	// 物流状态 揽收 发货 运输 签收等
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// 移动电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 发生时间
	OccurTime int64 `json:"occur_time,omitempty" xml:"occur_time,omitempty"`
	// 1-查件 2-寄件
	Scene int64 `json:"scene,omitempty" xml:"scene,omitempty"`
}

var poolTaowaiMsgSendRequest = sync.Pool{
	New: func() any {
		return new(TaowaiMsgSendRequest)
	},
}

// GetTaowaiMsgSendRequest() 从对象池中获取TaowaiMsgSendRequest
func GetTaowaiMsgSendRequest() *TaowaiMsgSendRequest {
	return poolTaowaiMsgSendRequest.Get().(*TaowaiMsgSendRequest)
}

// ReleaseTaowaiMsgSendRequest 释放TaowaiMsgSendRequest
func ReleaseTaowaiMsgSendRequest(v *TaowaiMsgSendRequest) {
	v.CpCode = ""
	v.MailNo = ""
	v.ExtParams = ""
	v.Openid = ""
	v.LogisticsStatus = ""
	v.MobilePhone = ""
	v.OccurTime = 0
	v.Scene = 0
	poolTaowaiMsgSendRequest.Put(v)
}
