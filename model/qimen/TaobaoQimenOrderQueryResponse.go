package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号接口 APIResponse
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。
*/
type TaobaoQimenOrderQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderQueryResponse
}

type TaobaoQimenOrderQueryResponse struct {
    XMLName xml.Name `xml:"qimen_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
