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
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_orderprocess_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *OrderProcessQueryResponse `json:"response,omitempty" xml:"