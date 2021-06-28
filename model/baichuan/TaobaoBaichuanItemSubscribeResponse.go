package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个商品订阅 APIResponse
taobao.baichuan.item.subscribe

百川单个商品订阅
*/
type TaobaoBaichuanItemSubscribeAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemSubscribeResponse
}

type TaobaoBaichuanItemSubscribeResponse struct {
    XMLName xml.Name `xml:"baichuan_item_subscribe_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemSubscribeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
