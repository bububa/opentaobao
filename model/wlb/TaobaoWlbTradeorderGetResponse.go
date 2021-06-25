package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据交易号获取物流宝订单 APIResponse
taobao.wlb.tradeorder.get

根据交易类型和交易id查询物流宝订单详情
*/
type TaobaoWlbTradeorderGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbTradeorderGetResponse `json:"taobao_wlb_tradeorder_get_response,omitempty"`
}

type TaobaoWlbTradeorderGetResponse struct {

    // 物流宝订单对象
    WlbOrderList   []WlbOrder `json:"wlb_order_list,omitempty"`

}
