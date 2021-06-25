package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益发放的查询接口 APIResponse
tmall.servicecenter.tp.funds.send.query

服务商资金权益发放结果的查询接口
*/
type TmallServicecenterTpFundsSendQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterTpFundsSendQueryResponse `json:"tmall_servicecenter_tp_funds_send_query_response,omitempty"`
}

type TmallServicecenterTpFundsSendQueryResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
