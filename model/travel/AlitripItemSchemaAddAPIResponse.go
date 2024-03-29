package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemSchemaAddAPIResponse 使用schema模板进行商品发布 API返回值
// alitrip.item.schema.add
//
// 飞猪度假商品使用schema模板进行商品发布。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemSchemaAddAPIResponse struct {
	model.CommonResponse
	AlitripItemSchemaAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripItemSchemaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripItemSchemaAddAPIResponseModel).Reset()
}

// AlitripItemSchemaAddAPIResponseModel is 使用schema模板进行商品发布 成功返回结果
type AlitripItemSchemaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_item_schema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopTravelItem `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripItemSchemaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripItemSchemaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripItemSchemaAddAPIResponse)
	},
}

// GetAlitripItemSchemaAddAPIResponse 从 sync.Pool 获取 AlitripItemSchemaAddAPIResponse
func GetAlitripItemSchemaAddAPIResponse() *AlitripItemSchemaAddAPIResponse {
	return poolAlitripItemSchemaAddAPIResponse.Get().(*AlitripItemSchemaAddAPIResponse)
}

// ReleaseAlitripItemSchemaAddAPIResponse 将 AlitripItemSchemaAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripItemSchemaAddAPIResponse(v *AlitripItemSchemaAddAPIResponse) {
	v.Reset()
	poolAlitripItemSchemaAddAPIResponse.Put(v)
}
