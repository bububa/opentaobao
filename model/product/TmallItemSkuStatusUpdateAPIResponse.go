package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskustatusupdateAPIResponse 商品sku状态更新 API返回值
// tmall.item.sku.status.update
//
// 商品sku上下架状态更新
type TmallitemskustatusupdateAPIResponse struct {
	model.CommonResponse
	TmallitemskustatusupdateAPIResponseModel
}

// TmallitemskustatusupdateAPIResponseModel is 商品sku状态更新 成功返回结果
type TmallitemskustatusupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallitemskustatusupdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
