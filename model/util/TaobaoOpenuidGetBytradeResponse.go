package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过订单获取对应买家的openUID APIResponse
taobao.openuid.get.bytrade

通过订单获取对应买家的openUID,需要卖家授权
*/
type TaobaoOpenuidGetBytradeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenuidGetBytradeResponse `json:"openuid_get_bytrade_response,omitempty"` 
    TaobaoOpenuidGetBytradeResponse
}

/* model for simplify = false
type TaobaoOpenuidGetBytradeResponse struct {

    // 当前交易tid对应买家的openuid
    
    OpenUid   string `json:"open_uid,omitempty"`
    

}
*/

type TaobaoOpenuidGetBytradeResponse struct {

    // 当前交易tid对应买家的openuid
    OpenUid   string `json:"open_uid,omitempty"`

}
