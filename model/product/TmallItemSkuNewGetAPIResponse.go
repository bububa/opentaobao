package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuNewGetAPIResponse 查询sku销售属性标新信息 API返回值
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
type TmallItemSkuNewGetAPIResponse struct {
	model.CommonResponse
	TmallItemSkuNewGetAPIResponseModel
}

// TmallItemSkuNewGetAPIResponseModel is 查询sku销售属性标新信息 成功返回结果
type TmallItemSkuNewGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_new_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallItemSkuNewGetApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
