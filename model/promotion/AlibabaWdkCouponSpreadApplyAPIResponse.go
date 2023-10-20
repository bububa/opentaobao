package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSpreadApplyAPIResponse 普通发券 API返回值
// alibaba.wdk.coupon.spread.apply
//
// 优惠券发放
type AlibabaWdkCouponSpreadApplyAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponSpreadApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponSpreadApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponSpreadApplyAPIResponseModel).Reset()
}

// AlibabaWdkCouponSpreadApplyAPIResponseModel is 普通发券 成功返回结果
type AlibabaWdkCouponSpreadApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_spread_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponSpreadApplyApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponSpreadApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponSpreadApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSpreadApplyAPIResponse)
	},
}

// GetAlibabaWdkCouponSpreadApplyAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponSpreadApplyAPIResponse
func GetAlibabaWdkCouponSpreadApplyAPIResponse() *AlibabaWdkCouponSpreadApplyAPIResponse {
	return poolAlibabaWdkCouponSpreadApplyAPIResponse.Get().(*AlibabaWdkCouponSpreadApplyAPIResponse)
}

// ReleaseAlibabaWdkCouponSpreadApplyAPIResponse 将 AlibabaWdkCouponSpreadApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponSpreadApplyAPIResponse(v *AlibabaWdkCouponSpreadApplyAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponSpreadApplyAPIResponse.Put(v)
}
