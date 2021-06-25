package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 APIResponse
taobao.vmarket.eticket.reverse

电子凭证平台冲正接口
*/
type TaobaoVmarketEticketReverseAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketReverseResponse `json:"taobao_vmarket_eticket_reverse_response,omitempty"`
}

type TaobaoVmarketEticketReverseResponse struct {

    // 0:失败，1:成功
    RetCode   int64 `json:"ret_code,omitempty"`

    // 宝贝标题
    ItemTitle   string `json:"item_title,omitempty"`

    // 整个订单的剩余可核销数量
    LeftNum   int64 `json:"left_num,omitempty"`

}
