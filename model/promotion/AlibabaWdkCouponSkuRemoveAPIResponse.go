package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponSkuRemoveAPIResponse 优惠券商品删除 API返回值
// alibaba.wdk.coupon.sku.remove
//
// 优惠券商品删除
type AlibabaWdkCouponSkuRemoveAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponSkuRemoveAPIResponseModel
}

// AlibabaWdkCouponSkuRemoveAPIResponseModel is 优惠券商品删除 成功返回结果
type AlibabaWdkCouponSkuRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponSkuRemoveApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
