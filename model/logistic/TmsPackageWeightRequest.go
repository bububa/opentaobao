package logistic

import (
	"sync"
)

// TmsPackageWeightRequest 结构体
type TmsPackageWeightRequest struct {
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 单据标识
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 包裹重量，单位克。若无重量，则需要通过abnormal_type字段给出原因
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 快递公司资源编码
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 无包裹重量或者包裹不准确的类型 0-包裹重量有效（默认有效） 1-无包裹重量 2-包裹重量不准确。若存在weight，则不需要传该值
	AbnormalType string `json:"abnormal_type,omitempty" xml:"abnormal_type,omitempty"`
	// 无包裹重量或者包裹不准确的原因。若存在weight，则不需要传该值
	AbnormalDesc string `json:"abnormal_desc,omitempty" xml:"abnormal_desc,omitempty"`
	// 服务商侧计费重量，单位g，小数点后2位
	ChargingWeight string `json:"charging_weight,omitempty" xml:"charging_weight,omitempty"`
	// 计抛重量，单位g，小数点后2位
	ThrowingWeight string `json:"throwing_weight,omitempty" xml:"throwing_weight,omitempty"`
	// 包裹长度 (单位：cm)，小数点后2位
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度 (单位：cm)，小数点后2位
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高度（单位：cm)，小数点后2位
	Height string `json:"height,omitempty" xml:"height,omitempty"`
}

var poolTmsPackageWeightRequest = sync.Pool{
	New: func() any {
		return new(TmsPackageWeightRequest)
	},
}

// GetTmsPackageWeightRequest() 从对象池中获取TmsPackageWeightRequest
func GetTmsPackageWeightRequest() *TmsPackageWeightRequest {
	return poolTmsPackageWeightRequest.Get().(*TmsPackageWeightRequest)
}

// ReleaseTmsPackageWeightRequest 释放TmsPackageWeightRequest
func ReleaseTmsPackageWeightRequest(v *TmsPackageWeightRequest) {
	v.MailNo = ""
	v.BizCode = ""
	v.Weight = ""
	v.TmsCpCode = ""
	v.AbnormalType = ""
	v.AbnormalDesc = ""
	v.ChargingWeight = ""
	v.ThrowingWeight = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	poolTmsPackageWeightRequest.Put(v)
}
