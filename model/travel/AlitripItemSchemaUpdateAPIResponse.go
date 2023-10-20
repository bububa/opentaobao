package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemSchemaUpdateAPIResponse 使用schema进行商品编辑 API返回值
// alitrip.item.schema.update
//
// 飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemSchemaUpdateAPIResponse struct {
	model.CommonResponse
	AlitripItemSchemaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripItemSchemaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripItemSchemaUpdateAPIResponseModel).Reset()
}

// AlitripItemSchemaUpdateAPIResponseModel is 使用schema进行商品编辑 成功返回结果
type AlitripItemSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_item_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopTravelItem `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripItemSchemaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripItemSchemaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripItemSchemaUpdateAPIResponse)
	},
}

// GetAlitripItemSchemaUpdateAPIResponse 从 sync.Pool 获取 AlitripItemSchemaUpdateAPIResponse
func GetAlitripItemSchemaUpdateAPIResponse() *AlitripItemSchemaUpdateAPIResponse {
	return poolAlitripItemSchemaUpdateAPIResponse.Get().(*AlitripItemSchemaUpdateAPIResponse)
}

// ReleaseAlitripItemSchemaUpdateAPIResponse 将 AlitripItemSchemaUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlitripItemSchemaUpdateAPIResponse(v *AlitripItemSchemaUpdateAPIResponse) {
	v.Reset()
	poolAlitripItemSchemaUpdateAPIResponse.Put(v)
}
