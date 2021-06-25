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
    Response *AlibabaWdkCouponSkuRemoveResponse `json:"alibaba_wdk_coupon_sku_remove_response,omitempty"`
}

type AlibabaWdkCouponSkuRemoveResponse struct {

    // 结果
    Result   *AlibabaWdkCouponSkuRemoveApiResult `json:"result,omitempty"`

}
