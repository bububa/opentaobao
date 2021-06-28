package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除商品订阅 APIResponse
taobao.baichuan.items.unsubscribe

批量删除商品订阅
*/
type TaobaoBaichuanItemsUnsubscribeAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemsUnsubscribeResponse
}

type TaobaoBaichuanItemsUnsubscribeResponse struct {
    XMLName xml.Name `xml:"baichuan_items_unsubscribe_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemsUnsubscribeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
