package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-新用户订单明细查询 APIResponse
taobao.tbk.dg.newuser.order.get

拉新API
*/
type TaobaoTbkDgNewuserOrderGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkDgNewuserOrderGetResponse `json:"taobao_tbk_dg_newuser_order_get_response,omitempty"`
}

type TaobaoTbkDgNewuserOrderGetResponse struct {

    // data
    Results   *Results `json:"results,omitempty"`

}
