package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemskunewupdateAPIResponse 更新sku销售属性标新状态 API返回值
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
type TmallitemskunewupdateAPIResponse struct {
	model.CommonResponse
	TmallitemskunewupdateAPIResponseModel
}

// TmallitemskunewupdateAPIResponseModel is 更新sku销售属性标新状态 成功返回结果
type TmallitemskunewupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_new_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallitemskunewupdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
