package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品删除 APIResponse
alibaba.wdk.coupon.sku.remove

优惠券商品删除
*/
type AlibabaWdkCouponSkuRemoveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkCouponSkuRemoveResponse `json:"alibaba_wdk_coupon_sku_remove_response,omitempty"` 
    AlibabaWdkCouponSkuRemoveResponse
}

/* model for simplify = false
type AlibabaWdkCouponSkuRemoveResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkCouponSkuRemoveApiResult  *AlibabaWdkCouponSkuRemoveApiResult `json:"alibaba_wdk_coupon_sku_remove_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkCouponSkuRemoveResponse struct {

    // 结果
    Result   *AlibabaWdkCouponSkuRemoveApiResult `json:"result,omitempty"`

}
