package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建物流订单 APIResponse
taobao.logistics.order.create

用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
*/
type TaobaoLogisticsOrderCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsOrderCreateResponse `json:"taobao_logistics_order_create_response,omitempty"`
}

type TaobaoLogisticsOrderCreateResponse struct {

    // 淘宝物流订单交易号，如返回-1则表示错误。如果在新建订单时传入trade_id,此处会返回此id，如果未传入trade_id，此处会返回淘宝物流分配的交易号码。
    TradeId   int64 `json:"trade_id,omitempty"`

}
