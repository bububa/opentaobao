package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品增加 APIResponse
alibaba.wdk.coupon.sku.add

优惠券商品增加
*/
type AlibabaWdkCouponSkuAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkCouponSkuAddResponse `json:"alibaba_wdk_coupon_sku_add_response,omitempty"` 
    AlibabaWdkCouponSkuAddResponse
}

/* model for simplify = false
type AlibabaWdkCouponSkuAddResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkCouponSkuAddApiResult  *AlibabaWdkCouponSkuAddApiResult `json:"alibaba_wdk_coupon_sku_add_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponSkuAddResponse struct {

    // 结果
    Result   *AlibabaWdkCouponSkuAddApiResult `json:"result,omitempty"`

}
