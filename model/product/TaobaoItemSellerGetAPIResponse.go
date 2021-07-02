package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSellerGetAPIResponse 获取单个商品详细信息 API返回值
// taobao.item.seller.get
//
// 获取单个商品的全部信息
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
type TaobaoItemSellerGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemSellerGetAPIResponseModel
}

// TaobaoItemSellerGetAPIResponseModel is 获取单个商品详细信息 成功返回结果
type TaobaoItemSellerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_seller_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品详细信息
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
