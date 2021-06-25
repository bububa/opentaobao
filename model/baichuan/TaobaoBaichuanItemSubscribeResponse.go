package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单个商品订阅 APIResponse
taobao.baichuan.item.subscribe

百川单个商品订阅
*/
type TaobaoBaichuanItemSubscribeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanItemSubscribeResponse `json:"taobao_baichuan_item_subscribe_response,omitempty"`
}

type TaobaoBaichuanItemSubscribeResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeResult `json:"result,omitempty"`

}
