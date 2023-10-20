package lstlogistics

import (
	"sync"
)

// SendOfflineOrderParam 结构体
type SendOfflineOrderParam struct {
	// 发货主订单列表
	MainOrderParamList []MainOrderParam `json:"main_order_param_list,omitempty" xml:"main_order_param_list>main_order_param,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 物流公司code
	CpCompanyCode string `json:"cp_company_code,omitempty" xml:"cp_company_code,omitempty"`
	// 物流公司名称
	CpCompanyName string `json:"cp_company_name,omitempty" xml:"cp_company_name,omitempty"`
	// 发货时间
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 备注
	Remarks string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

var poolSendOfflineOrderParam = sync.Pool{
	New: func() any {
		return new(SendOfflineOrderParam)
	},
}

// GetSendOfflineOrderParam() 从对象池中获取SendOfflineOrderParam
func GetSendOfflineOrderParam() *SendOfflineOrderParam {
	return poolSendOfflineOrderParam.Get().(*SendOfflineOrderParam)
}

// ReleaseSendOfflineOrderParam 释放SendOfflineOrderParam
func ReleaseSendOfflineOrderParam(v *SendOfflineOrderParam) {
	v.MainOrderParamList = v.MainOrderParamList[:0]
	v.MailNo = ""
	v.CpCompanyCode = ""
	v.CpCompanyName = ""
	v.SendTime = ""
	v.Remarks = ""
	poolSendOfflineOrderParam.Put(v)
}
