package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表 APIResponse
taobao.train.agent.bookorders.get

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentBookordersGetResponse `json:"taobao_train_agent_bookorders_get_response,omitempty"`
}

type TaobaoTrainAgentBookordersGetResponse struct {

    // 待处理订单总数
    OrderCount   int64 `json:"order_count,omitempty"`

    // 订单号集合，用半角逗号(,)连接，只会返回固定数量
    OrderIds   string `json:"order_ids,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

}
