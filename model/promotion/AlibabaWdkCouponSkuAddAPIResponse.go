package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSkuAddAPIResponse 优惠券商品增加 API返回值
// alibaba.wdk.coupon.sku.add
//
// 优惠券商品增加
type AlibabaWdkCouponSkuAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponSkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponSkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponSkuAddAPIResponseModel).Reset()
}

// AlibabaWdkCouponSkuAddAPIResponseModel is 优惠券商品增加 成功返回结果
type AlibabaWdkCouponSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponSkuAddApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponSkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponSkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponSkuAddAPIResponse)
	},
}

// GetAlibabaWdkCouponSkuAddAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponSkuAddAPIResponse
func GetAlibabaWdkCouponSkuAddAPIResponse() *AlibabaWdkCouponSkuAddAPIResponse {
	return poolAlibabaWdkCouponSkuAddAPIResponse.Get().(*AlibabaWdkCouponSkuAddAPIResponse)
}

// ReleaseAlibabaWdkCouponSkuAddAPIResponse 将 AlibabaWdkCouponSkuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponSkuAddAPIResponse(v *AlibabaWdkCouponSkuAddAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponSkuAddAPIResponse.Put(v)
}
