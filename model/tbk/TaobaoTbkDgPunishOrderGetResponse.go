package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-处罚订单查询 APIResponse
taobao.tbk.dg.punish.order.get

新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
*/
type TaobaoTbkDgPunishOrderGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkDgPunishOrderGetResponse `json:"taobao_tbk_dg_punish_order_get_response,omitempty"`
}

type TaobaoTbkDgPunishOrderGetResponse struct {

    // 查询的对象
    Result   *RpcResult `json:"result,omitempty"`

}