package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商品 API返回值 
alibaba.wdk.item.serviceitem.query

查询服务商品
*/
type AlibabaWdkItemServiceitemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemServiceitemQueryAPIResponseModel
}

// 查询服务商品 成功返回结果
type AlibabaWdkItemServiceitemQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_serviceitem_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkItemServiceitemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
