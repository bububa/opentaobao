package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelGereralproductUpdateAPIResponse 通用类目产品发布编辑 API返回值
// alitrip.travel.gereralproduct.update
//
// 提供给飞猪供销平台供应商发布编辑通用类目产品的API
type AlitripTravelGereralproductUpdateAPIResponse struct {
	model.CommonResponse
	AlitripTravelGereralproductUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelGereralproductUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelGereralproductUpdateAPIResponseModel).Reset()
}

// AlitripTravelGereralproductUpdateAPIResponseModel is 通用类目产品发布编辑 成功返回结果
type AlitripTravelGereralproductUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_gereralproduct_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// firstResult
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelGereralproductUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripTravelGereralproductUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelGereralproductUpdateAPIResponse)
	},
}

// GetAlitripTravelGereralproductUpdateAPIResponse 从 sync.Pool 获取 AlitripTravelGereralproductUpdateAPIResponse
func GetAlitripTravelGereralproductUpdateAPIResponse() *AlitripTravelGereralproductUpdateAPIResponse {
	return poolAlitripTravelGereralproductUpdateAPIResponse.Get().(*AlitripTravelGereralproductUpdateAPIResponse)
}

// ReleaseAlitripTravelGereralproductUpdateAPIResponse 将 AlitripTravelGereralproductUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelGereralproductUpdateAPIResponse(v *AlitripTravelGereralproductUpdateAPIResponse) {
	v.Reset()
	poolAlitripTravelGereralproductUpdateAPIResponse.Put(v)
}
