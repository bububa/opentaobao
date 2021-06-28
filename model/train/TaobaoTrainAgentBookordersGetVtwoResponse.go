package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表v2--增加鉴权校验 APIResponse
taobao.train.agent.bookorders.get.vtwo

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetVtwoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"train_agent_bookorders_get_vtwo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 待处理订单总数
    
    OrderCount   int64 `json:"order_count,omitempty" xml:"