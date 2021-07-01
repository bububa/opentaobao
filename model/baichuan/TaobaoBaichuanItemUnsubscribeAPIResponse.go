package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个删除订阅关系 API返回值 
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
type TaobaoBaichuanItemUnsubscribeAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemUnsubscribeAPIResponseModel
}

// 单个删除订阅关系 成功返回结果
type TaobaoBaichuanItemUnsubscribeAPIResponseModel struct {
    XMLName xml.Name `xml:"baichuan_item_unsubscribe_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoBaichuanItemUnsubscribeResult `json:"result,omitempty" xml:"result,omitempty"`
}
