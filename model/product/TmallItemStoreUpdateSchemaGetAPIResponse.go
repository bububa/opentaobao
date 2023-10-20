package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemStoreUpdateSchemaGetAPIResponse 天猫门店商品修改规则获取 API返回值
// tmall.item.store.update.schema.get
//
// 天猫门店商品修改规则获取
type TmallItemStoreUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemStoreUpdateSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemStoreUpdateSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemStoreUpdateSchemaGetAPIResponseModel).Reset()
}

// TmallItemStoreUpdateSchemaGetAPIResponseModel is 天猫门店商品修改规则获取 成功返回结果
type TmallItemStoreUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_store_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TmallItemStoreUpdateSchemaGetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemStoreUpdateSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTmallItemStoreUpdateSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemStoreUpdateSchemaGetAPIResponse)
	},
}

// GetTmallItemStoreUpdateSchemaGetAPIResponse 从 sync.Pool 获取 TmallItemStoreUpdateSchemaGetAPIResponse
func GetTmallItemStoreUpdateSchemaGetAPIResponse() *TmallItemStoreUpdateSchemaGetAPIResponse {
	return poolTmallItemStoreUpdateSchemaGetAPIResponse.Get().(*TmallItemStoreUpdateSchemaGetAPIResponse)
}

// ReleaseTmallItemStoreUpdateSchemaGetAPIResponse 将 TmallItemStoreUpdateSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemStoreUpdateSchemaGetAPIResponse(v *TmallItemStoreUpdateSchemaGetAPIResponse) {
	v.Reset()
	poolTmallItemStoreUpdateSchemaGetAPIResponse.Put(v)
}
