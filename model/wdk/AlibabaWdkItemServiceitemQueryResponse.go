package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商品 APIResponse
alibaba.wdk.item.serviceitem.query

查询服务商品
*/
type AlibabaWdkItemServiceitemQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkItemServiceitemQueryResponse `json:"alibaba_wdk_item_serviceitem_query_response,omitempty"`
}

type AlibabaWdkItemServiceitemQueryResponse struct {

    // result
    Result   *AlibabaWdkItemServiceitemQueryResult `json:"result,omitempty"`

}
