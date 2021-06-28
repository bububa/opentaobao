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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_coupon_sku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponSkuAddApiResult `json:"result,omitempty" xml:"