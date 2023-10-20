package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelProductGereralskuUpdateAPIResponse (供销)船票通用类目sku新增&编辑API API返回值
// alitrip.travel.product.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
type AlitripTravelProductGereralskuUpdateAPIResponse struct {
	model.CommonResponse
	AlitripTravelProductGereralskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelProductGereralskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelProductGereralskuUpdateAPIResponseModel).Reset()
}

// AlitripTravelProductGereralskuUpdateAPIResponseModel is (供销)船票通用类目sku新增&编辑API 成功返回结果
type AlitripTravelProductGereralskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_product_gereralsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelProductGereralskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripTravelProductGereralskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelProductGereralskuUpdateAPIResponse)
	},
}

// GetAlitripTravelProductGereralskuUpdateAPIResponse 从 sync.Pool 获取 AlitripTravelProductGereralskuUpdateAPIResponse
func GetAlitripTravelProductGereralskuUpdateAPIResponse() *AlitripTravelProductGereralskuUpdateAPIResponse {
	return poolAlitripTravelProductGereralskuUpdateAPIResponse.Get().(*AlitripTravelProductGereralskuUpdateAPIResponse)
}

// ReleaseAlitripTravelProductGereralskuUpdateAPIResponse 将 AlitripTravelProductGereralskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelProductGereralskuUpdateAPIResponse(v *AlitripTravelProductGereralskuUpdateAPIResponse) {
	v.Reset()
	poolAlitripTravelProductGereralskuUpdateAPIResponse.Put(v)
}
