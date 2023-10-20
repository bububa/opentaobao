package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemssubscribeAPIResponse 百川批量商品订阅 API返回值
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
type TaobaobaichuanitemssubscribeAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanitemssubscribeAPIResponseModel
}

// TaobaobaichuanitemssubscribeAPIResponseModel is 百川批量商品订阅 成功返回结果
type TaobaobaichuanitemssubscribeAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_items_subscribe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaobaichuanitemssubscribeResult `json:"result,omitempty" xml:"result,omitempty"`
}
