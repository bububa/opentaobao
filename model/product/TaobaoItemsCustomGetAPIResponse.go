package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemsCustomGetAPIResponse 根据外部ID取商品 API返回值
// taobao.items.custom.get
//
// 跟据卖家设定的商品外部id获取商品，只能获取授权卖家的商品
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
type TaobaoItemsCustomGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemsCustomGetAPIResponseModel
}

// TaobaoItemsCustomGetAPIResponseModel is 根据外部ID取商品 成功返回结果
type TaobaoItemsCustomGetAPIResponseModel struct {
	XMLName xml.Name `xml:"items_custom_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品列表，具体返回字段以fields决定
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
