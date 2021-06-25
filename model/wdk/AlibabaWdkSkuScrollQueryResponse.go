package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量游标方式查询接口 APIResponse
alibaba.wdk.sku.scroll.query

通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
*/
type AlibabaWdkSkuScrollQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuScrollQueryResponse `json:"alibaba_wdk_sku_scroll_query_response,omitempty"`
}

type AlibabaWdkSkuScrollQueryResponse struct {

    // 返回结果
    Result   *ApiScrollPageResult `json:"result,omitempty"`

}
