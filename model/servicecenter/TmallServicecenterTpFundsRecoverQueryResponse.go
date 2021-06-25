package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益逆向扣回的查询接口 APIResponse
tmall.servicecenter.tp.funds.recover.query

服务商资金权益逆向扣回的查询接口
*/
type TmallServicecenterTpFundsRecoverQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterTpFundsRecoverQueryResponse `json:"tmall_servicecenter_tp_funds_recover_query_response,omitempty"`
}

type TmallServicecenterTpFundsRecoverQueryResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
