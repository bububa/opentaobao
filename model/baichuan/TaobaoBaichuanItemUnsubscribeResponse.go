package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单个删除订阅关系 APIResponse
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
type TaobaoBaichuanItemUnsubscribeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanItemUnsubscribeResponse `json:"taobao_baichuan_item_unsubscribe_response,omitempty"`
}

type TaobaoBaichuanItemUnsubscribeResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemUnsubscribeResult `json:"result,omitempty"`

}
