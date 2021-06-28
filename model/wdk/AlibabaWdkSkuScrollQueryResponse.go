package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量游标方式查询接口 APIResponse
alibaba.wdk.sku.scroll.query

通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
*/
type AlibabaWdkSkuScrollQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuScrollQueryResponse
}

type AlibabaWdkSkuScrollQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_scroll_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
