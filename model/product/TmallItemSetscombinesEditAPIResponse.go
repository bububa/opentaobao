package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSetscombinesEditAPIResponse 普通商品转套装商品&套装商品编辑接口 API返回值
// tmall.item.setscombines.edit
//
// 普通商品转套装商品&套装商品编辑接口
type TmallItemSetscombinesEditAPIResponse struct {
	model.CommonResponse
	TmallItemSetscombinesEditAPIResponseModel
}

// TmallItemSetscombinesEditAPIResponseModel is 普通商品转套装商品&套装商品编辑接口 成功返回结果
type TmallItemSetscombinesEditAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_setscombines_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 套装货品更新结果
	SetscombineUpdateResult string `json:"setscombine_update_result,omitempty" xml:"setscombine_update_result,omitempty"`
}
