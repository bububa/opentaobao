package promotion

import (
	"encoding/xml"

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

// AlibabaWdkCouponSkuAddAPIResponseModel is 优惠券商品增加 成功返回结果
type AlibabaWdkCouponSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponSkuAddApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
