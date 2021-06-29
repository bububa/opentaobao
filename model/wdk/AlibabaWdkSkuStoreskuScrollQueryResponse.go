package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量查询接口 APIResponse
alibaba.wdk.sku.storesku.scroll.query

提供门店商品批量查询接口
*/
type AlibabaWdkSkuStoreskuScrollQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuStoreskuScrollQueryResponse
}

type AlibabaWdkSkuStoreskuScrollQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_storesku_scroll_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果
    
    Result   *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
