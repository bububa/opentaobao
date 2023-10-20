package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemstoreschemaupdateAPIResponse 天猫门店商品编辑 API返回值
// tmall.item.store.schema.update
//
// 天猫门店商品编辑
type TmallitemstoreschemaupdateAPIResponse struct {
	model.CommonResponse
	TmallitemstoreschemaupdateAPIResponseModel
}

// TmallitemstoreschemaupdateAPIResponseModel is 天猫门店商品编辑 成功返回结果
type TmallitemstoreschemaupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_store_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TmallitemstoreschemaupdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
