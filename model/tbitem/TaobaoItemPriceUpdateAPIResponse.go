package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPriceUpdateAPIResponse 更新商品价格 API返回值
// taobao.item.price.update
//
// 更新商品价格
type TaobaoItemPriceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoItemPriceUpdateAPIResponseModel
}

// TaobaoItemPriceUpdateAPIResponseModel is 更新商品价格 成功返回结果
type TaobaoItemPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"item_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品结构里的num_iid，modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
