package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益发放的查询接口 APIResponse
tmall.servicecenter.tp.funds.send.query

服务商资金权益发放结果的查询接口
*/
type TmallServicecenterTpFundsSendQueryAPIResponse struct {
    model.CommonResponse
    TmallServicecenterTpFundsSendQueryResponse
}

type TmallServicecenterTpFundsSendQueryResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_tp_funds_send_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
