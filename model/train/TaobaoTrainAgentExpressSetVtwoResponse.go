package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息v2--增加鉴权校验 APIResponse
taobao.train.agent.express.set.vtwo

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_express_set_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ErrorMsgCode   string `json:"error_msg_code,omitempty" xml:"