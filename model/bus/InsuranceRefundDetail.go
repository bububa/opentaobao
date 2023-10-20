package bus

import (
	"sync"
)

// InsuranceRefundDetail 结构体
type InsuranceRefundDetail struct {
	// 退保信息
	TvmInsuranceInfos []TvmInsuranceInfo `json:"tvm_insurance_infos,omitempty" xml:"tvm_insurance_infos>tvm_insurance_info,omitempty"`
	// 乘客证件号
	RiderCertNumber string `json:"rider_cert_number,omitempty" xml:"rider_cert_number,omitempty"`
	// 证件类型：01：身份证，02：护照，03：港澳通行证，04：台湾通行证，05：港澳往返内地通行证，06：台湾往返内地通行证，07：港澳居民居住证
	RiderCertType string `json:"rider_cert_type,omitempty" xml:"rider_cert_type,omitempty"`
	// 乘客姓名
	RiderName string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// 保险退款总金额
	InsurePrice int64 `json:"insure_price,omitempty" xml:"insure_price,omitempty"`
}

var poolInsuranceRefundDetail = sync.Pool{
	New: func() any {
		return new(InsuranceRefundDetail)
	},
}

// GetInsuranceRefundDetail() 从对象池中获取InsuranceRefundDetail
func GetInsuranceRefundDetail() *InsuranceRefundDetail {
	return poolInsuranceRefundDetail.Get().(*InsuranceRefundDetail)
}

// ReleaseInsuranceRefundDetail 释放InsuranceRefundDetail
func ReleaseInsuranceRefundDetail(v *InsuranceRefundDetail) {
	v.TvmInsuranceInfos = v.TvmInsuranceInfos[:0]
	v.RiderCertNumber = ""
	v.RiderCertType = ""
	v.RiderName = ""
	v.InsurePrice = 0
	poolInsuranceRefundDetail.Put(v)
}
