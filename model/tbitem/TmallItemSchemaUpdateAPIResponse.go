package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSchemaUpdateAPIResponse 天猫根据规则编辑商品 API返回值
// tmall.item.schema.update
//
// 天猫根据规则编辑商品
type TmallItemSchemaUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSchemaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSchemaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSchemaUpdateAPIResponseModel).Reset()
}

// TmallItemSchemaUpdateAPIResponseModel is 天猫根据规则编辑商品 成功返回结果
type TmallItemSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回商品发布结果
	UpdateItemResult string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`
	// 商品更新操作成功时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSchemaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateItemResult = ""
	m.GmtModified = ""
}

var poolTmallItemSchemaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSchemaUpdateAPIResponse)
	},
}

// GetTmallItemSchemaUpdateAPIResponse 从 sync.Pool 获取 TmallItemSchemaUpdateAPIResponse
func GetTmallItemSchemaUpdateAPIResponse() *TmallItemSchemaUpdateAPIResponse {
	return poolTmallItemSchemaUpdateAPIResponse.Get().(*TmallItemSchemaUpdateAPIResponse)
}

// ReleaseTmallItemSchemaUpdateAPIResponse 将 TmallItemSchemaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSchemaUpdateAPIResponse(v *TmallItemSchemaUpdateAPIResponse) {
	v.Reset()
	poolTmallItemSchemaUpdateAPIResponse.Put(v)
}
