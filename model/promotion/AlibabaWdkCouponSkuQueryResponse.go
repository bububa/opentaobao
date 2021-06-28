package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品查询 APIResponse
alibaba.wdk.coupon.sku.query

优惠券商品查询
*/
type AlibabaWdkCouponSkuQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_coupon_sku_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponSkuQueryApiResult `json:"result,omitempty" xml:"