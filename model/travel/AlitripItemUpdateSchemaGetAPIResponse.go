package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemUpdateSchemaGetAPIResponse 获取编辑商品的schema模板 API返回值
// alitrip.item.update.schema.get
//
// 获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	AlitripItemUpdateSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripItemUpdateSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripItemUpdateSchemaGetAPIResponseModel).Reset()
}

// AlitripItemUpdateSchemaGetAPIResponseModel is 获取编辑商品的schema模板 成功返回结果
type AlitripItemUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_item_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// schema数据
	SchemaXmlFields string `json:"schema_xml_fields,omitempty" xml:"schema_xml_fields,omitempty"`
}

// Reset 清空结构体
func (m *AlitripItemUpdateSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SchemaXmlFields = ""
}

var poolAlitripItemUpdateSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripItemUpdateSchemaGetAPIResponse)
	},
}

// GetAlitripItemUpdateSchemaGetAPIResponse 从 sync.Pool 获取 AlitripItemUpdateSchemaGetAPIResponse
func GetAlitripItemUpdateSchemaGetAPIResponse() *AlitripItemUpdateSchemaGetAPIResponse {
	return poolAlitripItemUpdateSchemaGetAPIResponse.Get().(*AlitripItemUpdateSchemaGetAPIResponse)
}

// ReleaseAlitripItemUpdateSchemaGetAPIResponse 将 AlitripItemUpdateSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripItemUpdateSchemaGetAPIResponse(v *AlitripItemUpdateSchemaGetAPIResponse) {
	v.Reset()
	poolAlitripItemUpdateSchemaGetAPIResponse.Put(v)
}
