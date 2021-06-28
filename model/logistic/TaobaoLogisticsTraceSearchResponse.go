package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流流转信息查询 APIResponse
taobao.logistics.trace.search

用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。<br/>此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
*/
type TaobaoLogisticsTraceSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"logistics_trace_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 运单号
    
    OutSid   string `json:"out_sid,omitempty" xml:"