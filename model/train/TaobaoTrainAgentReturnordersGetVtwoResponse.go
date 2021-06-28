package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待退票的订单v2--增加鉴权校验 APIResponse
taobao.train.agent.returnorders.get.vtwo

代理商用来获取待退票的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentReturnordersGetVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_returnorders_get_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 待退票的订单数
    
    OrderCount   int64 `json:"order_count,omitempty" xml:"