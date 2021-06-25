package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川批量商品订阅 APIResponse
taobao.baichuan.items.subscribe

百川批量添加订阅的商品
*/
type TaobaoBaichuanItemsSubscribeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanItemsSubscribeResponse `json:"taobao_baichuan_items_subscribe_response,omitempty"`
}

type TaobaoBaichuanItemsSubscribeResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemsSubscribeResult `json:"result,omitempty"`

}
