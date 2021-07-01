package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemPriceUpdateAPIResponse
天猫商品/SKU价格更新接口 API返回值
tmall.item.price.update

天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。 */
type TmallItemPriceUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemPriceUpdateAPIResponseModel
}

// TmallItemPriceUpdateAPIResponseModel is 天猫商品/SKU价格更新接口 成功返回结果
type TmallItemPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 价格更新结果
	PriceUpdateResult string `json:"price_update_result,omitempty" xml:"price_update_result,omitempty"`
}
