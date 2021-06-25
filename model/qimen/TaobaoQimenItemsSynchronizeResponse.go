package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 (批量) APIResponse
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS
*/
type TaobaoQimenItemsSynchronizeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenItemsSynchronizeResponse `json:"taobao_qimen_items_synchronize_response,omitempty"`
}

type TaobaoQimenItemsSynchronizeResponse struct {

    // 
    Response   *ItemsSynchronizeResponse `json:"response,omitempty"`

}
