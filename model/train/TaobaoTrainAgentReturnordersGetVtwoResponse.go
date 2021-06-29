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
    TaobaoTrainAgentReturnordersGetVtwoResponse
}

type TaobaoTrainAgentReturnordersGetVtwoResponse struct {
    XMLName xml.Name `xml:"train_agent_returnorders_get_vtwo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 待退票的订单数
    
    OrderCount   int64 `json:"order_count,omitempty" xml:"order_count,omitempty"`

    
    // 子订单号字符串，每个订单以逗号分隔
    
    OrderIds   string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`

    
    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 主订单id列表
    
    MainOrderIds   string `json:"main_order_ids,omitempty" xml:"main_order_ids,omitempty"`

    
    // 申请时间，每个时间以逗号分隔
    
    RefundApplyTimes   string `json:"refund_apply_times,omitempty" xml:"refund_apply_times,omitempty"`

    
}
