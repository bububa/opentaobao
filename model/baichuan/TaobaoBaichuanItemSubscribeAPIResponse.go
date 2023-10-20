package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsubscribeAPIResponse 单个商品订阅 API返回值
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
type TaobaobaichuanitemsubscribeAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanitemsubscribeAPIResponseModel
}

// TaobaobaichuanitemsubscribeAPIResponseModel is 单个商品订阅 成功返回结果
type TaobaobaichuanitemsubscribeAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_item_subscribe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaobaichuanitemsubscribeResult `json:"result,omitempty" xml:"result,omitempty"`
}
