package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】订单搜索 API返回值 
taobao.jipiao.agent.order.search

卖家根据条件查询淘宝订单id列表
*/
type TaobaoJipiaoAgentOrderSearchAPIResponse struct {
    model.CommonResponse
    TaobaoJipiaoAgentOrderSearchResponse
}

// 【机票代理商订单】订单搜索 成功返回结果
type TaobaoJipiaoAgentOrderSearchResponse struct {
    XMLName xml.Name `xml:"jipiao_agent_order_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回操作成功失败信息
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 失败时的错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 机票订单搜索返回结果对象
    SearchResult   *SearchOrderResult `json:"search_result,omitempty" xml:"search_result,omitempty"`
}
