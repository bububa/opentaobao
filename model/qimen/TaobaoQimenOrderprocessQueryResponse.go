package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单流水查询接口 APIResponse
taobao.qimen.orderprocess.query

ERP调用订单流水查询接口
*/
type TaobaoQimenOrderprocessQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderprocessQueryResponse
}

type TaobaoQimenOrderprocessQueryResponse struct {
    XMLName xml.Name `xml:"qimen_orderprocess_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *OrderProcessQueryResponse `json:"response,omitempty" xml:"response,omitempty"`

    
}
