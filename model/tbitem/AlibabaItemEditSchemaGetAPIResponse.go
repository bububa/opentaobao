package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemEditSchemaGetAPIResponse 商品编辑获取schema信息 API返回值
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
type AlibabaItemEditSchemaGetAPIResponse struct {
	model.CommonResponse
	AlibabaItemEditSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemEditSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemEditSchemaGetAPIResponseModel).Reset()
}

// AlibabaItemEditSchemaGetAPIResponseModel is 商品编辑获取schema信息 成功返回结果
type AlibabaItemEditSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_edit_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品已有规则信息，XML格式.
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemEditSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItemEditSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemEditSchemaGetAPIResponse)
	},
}

// GetAlibabaItemEditSchemaGetAPIResponse 从 sync.Pool 获取 AlibabaItemEditSchemaGetAPIResponse
func GetAlibabaItemEditSchemaGetAPIResponse() *AlibabaItemEditSchemaGetAPIResponse {
	return poolAlibabaItemEditSchemaGetAPIResponse.Get().(*AlibabaItemEditSchemaGetAPIResponse)
}

// ReleaseAlibabaItemEditSchemaGetAPIResponse 将 AlibabaItemEditSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemEditSchemaGetAPIResponse(v *AlibabaItemEditSchemaGetAPIResponse) {
	v.Reset()
	poolAlibabaItemEditSchemaGetAPIResponse.Put(v)
}
