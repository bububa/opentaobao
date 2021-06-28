package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息 APIResponse
taobao.train.agent.express.set

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_express_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ErrorMsgCode   string `json:"error_msg_code,omitempty" xml:"