package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskusortupdateAPIResponse 商品销售属性排序更新 API返回值
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
type TmallitemskusortupdateAPIResponse struct {
	model.CommonResponse
	TmallitemskusortupdateAPIResponseModel
}

// TmallitemskusortupdateAPIResponseModel is 商品销售属性排序更新 成功返回结果
type TmallitemskusortupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_sort_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallitemskusortupdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
