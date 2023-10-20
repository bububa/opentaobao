package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemUpdateSchemaGetAPIResponse 天猫编辑商品规则获取 API返回值
// tmall.item.update.schema.get
//
// Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemUpdateSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemUpdateSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemUpdateSchemaGetAPIResponseModel).Reset()
}

// TmallItemUpdateSchemaGetAPIResponseModel is 天猫编辑商品规则获取 成功返回结果
type TmallItemUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回发布商品的规则文档
	UpdateItemResult string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemUpdateSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateItemResult = ""
}

var poolTmallItemUpdateSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemUpdateSchemaGetAPIResponse)
	},
}

// GetTmallItemUpdateSchemaGetAPIResponse 从 sync.Pool 获取 TmallItemUpdateSchemaGetAPIResponse
func GetTmallItemUpdateSchemaGetAPIResponse() *TmallItemUpdateSchemaGetAPIResponse {
	return poolTmallItemUpdateSchemaGetAPIResponse.Get().(*TmallItemUpdateSchemaGetAPIResponse)
}

// ReleaseTmallItemUpdateSchemaGetAPIResponse 将 TmallItemUpdateSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemUpdateSchemaGetAPIResponse(v *TmallItemUpdateSchemaGetAPIResponse) {
	v.Reset()
	poolTmallItemUpdateSchemaGetAPIResponse.Put(v)
}
