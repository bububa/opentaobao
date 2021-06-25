package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据条件删除订阅关系 APIResponse
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系
*/
type TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanItemsUnsubscribeByConditionResponse `json:"taobao_baichuan_items_unsubscribe_by_condition_response,omitempty"`
}

type TaobaoBaichuanItemsUnsubscribeByConditionResponse struct {

    // 接口返回model
    Result   *TaobaoBaichuanItemsUnsubscribeByConditionResult `json:"result,omitempty"`

}
