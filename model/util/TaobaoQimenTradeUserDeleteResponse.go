package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除奇门订单链路用户 APIResponse
taobao.qimen.trade.user.delete

删除奇门订单链路用户
*/
type TaobaoQimenTradeUserDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenTradeUserDeleteResponse `json:"taobao_qimen_trade_user_delete_response,omitempty"`
}

type TaobaoQimenTradeUserDeleteResponse struct {

    // modal
    Modal   bool `json:"modal,omitempty"`

}
