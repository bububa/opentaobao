package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
店铺同步接口 APIResponse
taobao.qimen.shop.synchronize

店铺同步接口描述
*/
type TaobaoQimenShopSynchronizeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenShopSynchronizeResponse `json:"taobao_qimen_shop_synchronize_response,omitempty"`
}

type TaobaoQimenShopSynchronizeResponse struct {

    // Response
    Response   *Response `json:"response,omitempty"`

}
