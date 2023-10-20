package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDataCouponGetAPIResponse 获取优惠券信息 API返回值
// alibaba.data.coupon.get
//
// 获取优惠券信息，仅作客户端鉴权虚拟api使用
type AlibabaDataCouponGetAPIResponse struct {
	model.CommonResponse
	AlibabaDataCouponGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDataCouponGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDataCouponGetAPIResponseModel).Reset()
}

// AlibabaDataCouponGetAPIResponseModel is 获取优惠券信息 成功返回结果
type AlibabaDataCouponGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_data_coupon_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// unnamed
	Unnamed string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDataCouponGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Unnamed = ""
}

var poolAlibabaDataCouponGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDataCouponGetAPIResponse)
	},
}

// GetAlibabaDataCouponGetAPIResponse 从 sync.Pool 获取 AlibabaDataCouponGetAPIResponse
func GetAlibabaDataCouponGetAPIResponse() *AlibabaDataCouponGetAPIResponse {
	return poolAlibabaDataCouponGetAPIResponse.Get().(*AlibabaDataCouponGetAPIResponse)
}

// ReleaseAlibabaDataCouponGetAPIResponse 将 AlibabaDataCouponGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDataCouponGetAPIResponse(v *AlibabaDataCouponGetAPIResponse) {
	v.Reset()
	poolAlibabaDataCouponGetAPIResponse.Put(v)
}
