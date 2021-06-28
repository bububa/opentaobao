package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品增加 APIResponse
alibaba.wdk.coupon.sku.add

优惠券商品增加
*/
type AlibabaWdkCouponSkuAddAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponSkuAddResponse
}

type AlibabaWdkCouponSkuAddResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponSkuAddApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
