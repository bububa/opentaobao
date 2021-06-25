package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品查询 APIResponse
alibaba.wdk.coupon.sku.query

优惠券商品查询
*/
type AlibabaWdkCouponSkuQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkCouponSkuQueryResponse `json:"alibaba_wdk_coupon_sku_query_response,omitempty"`
}

type AlibabaWdkCouponSkuQueryResponse struct {

    // 结果
    Result   *AlibabaWdkCouponSkuQueryApiResult `json:"result,omitempty"`

}
