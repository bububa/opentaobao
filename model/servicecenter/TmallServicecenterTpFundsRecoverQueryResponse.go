package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商资金权益逆向扣回的查询接口 APIResponse
tmall.servicecenter.tp.funds.recover.query

服务商资金权益逆向扣回的查询接口
*/
type TmallServicecenterTpFundsRecoverQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_tp_funds_recover_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"