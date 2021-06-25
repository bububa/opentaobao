package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量删除商品订阅 APIResponse
taobao.baichuan.items.unsubscribe

批量删除商品订阅
*/
type TaobaoBaichuanItemsUnsubscribeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanItemsUnsubscribeResponse `json:"taobao_baichuan_items_unsubscribe_response,omitempty"`
}

type TaobaoBaichuanItemsUnsubscribeResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemsUnsubscribeResult `json:"result,omitempty"`

}
