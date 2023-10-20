package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemStoreSchemaUpdateAPIResponse 天猫门店商品编辑 API返回值
// tmall.item.store.schema.update
//
// 天猫门店商品编辑
type TmallItemStoreSchemaUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemStoreSchemaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemStoreSchemaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemStoreSchemaUpdateAPIResponseModel).Reset()
}

// TmallItemStoreSchemaUpdateAPIResponseModel is 天猫门店商品编辑 成功返回结果
type TmallItemStoreSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_store_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TmallItemStoreSchemaUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemStoreSchemaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTmallItemStoreSchemaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemStoreSchemaUpdateAPIResponse)
	},
}

// GetTmallItemStoreSchemaUpdateAPIResponse 从 sync.Pool 获取 TmallItemStoreSchemaUpdateAPIResponse
func GetTmallItemStoreSchemaUpdateAPIResponse() *TmallItemStoreSchemaUpdateAPIResponse {
	return poolTmallItemStoreSchemaUpdateAPIResponse.Get().(*TmallItemStoreSchemaUpdateAPIResponse)
}

// ReleaseTmallItemStoreSchemaUpdateAPIResponse 将 TmallItemStoreSchemaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemStoreSchemaUpdateAPIResponse(v *TmallItemStoreSchemaUpdateAPIResponse) {
	v.Reset()
	poolTmallItemStoreSchemaUpdateAPIResponse.Put(v)
}
