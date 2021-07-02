package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkusGetAPIResponse 根据商品ID列表获取SKU信息 API返回值
// taobao.item.skus.get
//
// * 获取多个商品下的所有sku
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
type TaobaoItemSkusGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemSkusGetAPIResponseModel
}

// TaobaoItemSkusGetAPIResponseModel is 根据商品ID列表获取SKU信息 成功返回结果
type TaobaoItemSkusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_skus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Sku列表
	Skus []Sku `json:"skus,omitempty" xml:"skus>sku,omitempty"`
}
