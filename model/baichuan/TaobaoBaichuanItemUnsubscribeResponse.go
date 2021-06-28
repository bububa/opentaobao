package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个删除订阅关系 APIResponse
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
type TaobaoBaichuanItemUnsubscribeAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemUnsubscribeResponse
}

type TaobaoBaichuanItemUnsubscribeResponse struct {
    XMLName xml.Name `xml:"baichuan_item_unsubscribe_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemUnsubscribeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
