package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退票通知 APIResponse
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调
*/
type TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_returnticket_confirm_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"