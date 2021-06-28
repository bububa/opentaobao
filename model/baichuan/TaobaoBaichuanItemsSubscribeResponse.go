package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川批量商品订阅 APIResponse
taobao.baichuan.items.subscribe

百川批量添加订阅的商品
*/
type TaobaoBaichuanItemsSubscribeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"baichuan_items_subscribe_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemsSubscribeResult `json:"result,omitempty" xml:"