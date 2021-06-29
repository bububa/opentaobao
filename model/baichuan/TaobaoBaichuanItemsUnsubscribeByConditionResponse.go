package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据条件删除订阅关系 API返回值 
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系
*/
type TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemsUnsubscribeByConditionResponse
}

// 根据条件删除订阅关系 成功返回结果
type TaobaoBaichuanItemsUnsubscribeByConditionResponse struct {
    XMLName xml.Name `xml:"baichuan_items_unsubscribe_by_condition_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoBaichuanItemsUnsubscribeByConditionResult `json:"result,omitempty" xml:"result,omitempty"`
}
