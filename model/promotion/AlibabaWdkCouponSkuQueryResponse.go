package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券商品查询 API返回值 
alibaba.wdk.coupon.sku.query

优惠券商品查询
*/
type AlibabaWdkCouponSkuQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponSkuQueryResponse
}

// 优惠券商品查询 成功返回结果
type AlibabaWdkCouponSkuQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaWdkCouponSkuQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
