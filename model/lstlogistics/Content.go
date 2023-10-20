package lstlogistics

import (
	"sync"
)

// Content 结构体
type Content struct {
	// 子订单列表
	SubOrderIdList []int64 `json:"sub_order_id_list,omitempty" xml:"sub_order_id_list>int64,omitempty"`
	// 物流编号，可用来查询物流详情
	LogisticsId string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
	// 运单编号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 物流公司code
	CpCompanyCode string `json:"cp_company_code,omitempty" xml:"cp_company_code,omitempty"`
	// 物流公司name
	CpCompanyName string `json:"cp_company_name,omitempty" xml:"cp_company_name,omitempty"`
	// 描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 描述
	StanderdDesc string `json:"standerd_desc,omitempty" xml:"standerd_desc,omitempty"`
}

var poolContent = sync.Pool{
	New: func() any {
		return new(Content)
	},
}

// GetContent() 从对象池中获取Content
func GetContent() *Content {
	return poolContent.Get().(*Content)
}

// ReleaseContent 释放Content
func ReleaseContent(v *Content) {
	v.SubOrderIdList = v.SubOrderIdList[:0]
	v.LogisticsId = ""
	v.MailNo = ""
	v.CpCompanyCode = ""
	v.CpCompanyName = ""
	v.StatusDesc = ""
	v.Time = ""
	v.StanderdDesc = ""
	poolContent.Put(v)
}
