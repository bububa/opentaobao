package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponAbandonAPIResponse 废券 API返回值
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
type AlibabaWdkCouponAbandonAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponAbandonAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponAbandonAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponAbandonAPIResponseModel).Reset()
}

// AlibabaWdkCouponAbandonAPIResponseModel is 废券 成功返回结果
type AlibabaWdkCouponAbandonAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_abandon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponAbandonApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponAbandonAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponAbandonAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponAbandonAPIResponse)
	},
}

// GetAlibabaWdkCouponAbandonAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponAbandonAPIResponse
func GetAlibabaWdkCouponAbandonAPIResponse() *AlibabaWdkCouponAbandonAPIResponse {
	return poolAlibabaWdkCouponAbandonAPIResponse.Get().(*AlibabaWdkCouponAbandonAPIResponse)
}

// ReleaseAlibabaWdkCouponAbandonAPIResponse 将 AlibabaWdkCouponAbandonAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponAbandonAPIResponse(v *AlibabaWdkCouponAbandonAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponAbandonAPIResponse.Put(v)
}
