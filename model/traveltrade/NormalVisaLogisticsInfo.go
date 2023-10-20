package traveltrade

import (
	"sync"
)

// NormalVisaLogisticsInfo 结构体
type NormalVisaLogisticsInfo struct {
	// 必填，物流号
	PostNumber string `json:"post_number,omitempty" xml:"post_number,omitempty"`
	// 必填，物流公司编码
	PostCompanyCode string `json:"post_company_code,omitempty" xml:"post_company_code,omitempty"`
	// 必填，物流公司名称
	PostCompanyName string `json:"post_company_name,omitempty" xml:"post_company_name,omitempty"`
	// 选填，物流联系人手机号(顺丰物流需要)
	ConcatPhone string `json:"concat_phone,omitempty" xml:"concat_phone,omitempty"`
}

var poolNormalVisaLogisticsInfo = sync.Pool{
	New: func() any {
		return new(NormalVisaLogisticsInfo)
	},
}

// GetNormalVisaLogisticsInfo() 从对象池中获取NormalVisaLogisticsInfo
func GetNormalVisaLogisticsInfo() *NormalVisaLogisticsInfo {
	return poolNormalVisaLogisticsInfo.Get().(*NormalVisaLogisticsInfo)
}

// ReleaseNormalVisaLogisticsInfo 释放NormalVisaLogisticsInfo
func ReleaseNormalVisaLogisticsInfo(v *NormalVisaLogisticsInfo) {
	v.PostNumber = ""
	v.PostCompanyCode = ""
	v.PostCompanyName = ""
	v.ConcatPhone = ""
	poolNormalVisaLogisticsInfo.Put(v)
}
