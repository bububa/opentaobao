package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商品 APIResponse
alibaba.wdk.item.serviceitem.query

查询服务商品
*/
type AlibabaWdkItemServiceitemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemServiceitemQueryResponse
}

type AlibabaWdkItemServiceitemQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_serviceitem_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkItemServiceitemQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
