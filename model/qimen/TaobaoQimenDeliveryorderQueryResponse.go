package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单查询接口 APIResponse
taobao.qimen.deliveryorder.query

ERP调用奇门的发货单查询接口，查询发货单详情
*/
type TaobaoQimenDeliveryorderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_deliveryorder_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *DeliveryOrderQueryResponse `json:"response,omitempty" xml:"