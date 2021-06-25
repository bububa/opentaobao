package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单记录导出 APIResponse
taobao.vas.order.search

用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。<br/>建议用于查询前一日的历史记录，不适合用作实时数据查询。<br/>现在只能查询90天以内的数据<br/>该接口限制每分钟所有appkey调用总和只能有800次。
*/
type TaobaoVasOrderSearchAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVasOrderSearchResponse `json:"taobao_vas_order_search_response,omitempty"`
}

type TaobaoVasOrderSearchResponse struct {

    // 商品订单对象
    ArticleBizOrders   []ArticleBizOrder `json:"article_biz_orders,omitempty"`

    // 总记录数
    TotalItem   int64 `json:"total_item,omitempty"`

}
