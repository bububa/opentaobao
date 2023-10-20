package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSkuQueryAPIResponse 优惠券商品查询 API返回值
// alibaba.wdk.coupon.sku.query
//
// 优惠券商品查询
type AlibabaWdkCouponSkuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponSkuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponSkuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponSkuQueryAPIResponseModel).Reset()
}

// AlibabaWdkCouponSkuQueryAPIResponseModel is 优惠券商品查询 成功返回结果
type AlibabaWdkCouponSkuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponSkuQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponSkuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponSkuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSkuQueryAPIResponse)
	},
}

// GetAlibabaWdkCouponSkuQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponSkuQueryAPIResponse
func GetAlibabaWdkCouponSkuQueryAPIResponse() *AlibabaWdkCouponSkuQueryAPIResponse {
	return poolAlibabaWdkCouponSkuQueryAPIResponse.Get().(*AlibabaWdkCouponSkuQueryAPIResponse)
}

// ReleaseAlibabaWdkCouponSkuQueryAPIResponse 将 AlibabaWdkCouponSkuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponSkuQueryAPIResponse(v *AlibabaWdkCouponSkuQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponSkuQueryAPIResponse.Put(v)
}
