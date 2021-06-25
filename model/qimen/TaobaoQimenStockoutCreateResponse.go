package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出库单创建接口 APIResponse
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息
*/
type TaobaoQimenStockoutCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenStockoutCreateResponse `json:"taobao_qimen_stockout_create_response,omitempty"`
}

type TaobaoQimenStockoutCreateResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
