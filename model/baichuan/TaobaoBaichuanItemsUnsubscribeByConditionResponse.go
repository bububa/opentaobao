package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据条件删除订阅关系 APIResponse
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系
*/
type TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"baichuan_items_unsubscribe_by_condition_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemsUnsubscribeByConditionResult `json:"result,omitempty" xml:"