package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单延时接口 APIResponse
taobao.vmarket.eticket.time.expand

提供码商操作订单延期接口
*/
type TaobaoVmarketEticketTimeExpandAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketTimeExpandResponse `json:"taobao_vmarket_eticket_time_expand_response,omitempty"`
}

type TaobaoVmarketEticketTimeExpandResponse struct {

    // 0:失败；1:成功
    RetCode   int64 `json:"ret_code,omitempty"`

}
