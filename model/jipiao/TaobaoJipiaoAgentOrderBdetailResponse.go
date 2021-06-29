package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】采购订单详情 APIResponse
taobao.jipiao.agent.order.bdetail

根据淘宝系统订单号获取订单详情信息
*/
type TaobaoJipiaoAgentOrderBdetailAPIResponse struct {
    model.CommonResponse
    TaobaoJipiaoAgentOrderBdetailResponse
}

type TaobaoJipiaoAgentOrderBdetailResponse struct {
    XMLName xml.Name `xml:"jipiao_agent_order_bdetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回操作成功失败信息
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 机票订单的详情列表，当前支持返回一个订单
    
    Orders   []TripOrder `json:"orders,omitempty" xml:"orders>trip_order,omitempty"`
    
    
}
