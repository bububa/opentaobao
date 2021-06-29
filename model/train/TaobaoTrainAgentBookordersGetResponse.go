package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表 APIResponse
taobao.train.agent.bookorders.get

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentBookordersGetResponse
}

type TaobaoTrainAgentBookordersGetResponse struct {
    XMLName xml.Name `xml:"train_agent_bookorders_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 待处理订单总数
    
    OrderCount   int64 `json:"order_count,omitempty" xml:"order_count,omitempty"`

    
    // 订单号集合，用半角逗号(,)连接，只会返回固定数量
    
    OrderIds   string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`

    
    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
