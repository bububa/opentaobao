package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCityCarApplyQueryAPIResponse 三方市内用车申请单查询 API返回值
// alitrip.btrip.city.car.apply.query
//
// 三方市内用车申请单查询
type AlitripBtripCityCarApplyQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripCityCarApplyQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCityCarApplyQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCityCarApplyQueryAPIResponseModel).Reset()
}

// AlitripBtripCityCarApplyQueryAPIResponseModel is 三方市内用车申请单查询 成功返回结果
type AlitripBtripCityCarApplyQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_city_car_apply_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCityCarApplyQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCityCarApplyQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCityCarApplyQueryAPIResponse)
	},
}

// GetAlitripBtripCityCarApplyQueryAPIResponse 从 sync.Pool 获取 AlitripBtripCityCarApplyQueryAPIResponse
func GetAlitripBtripCityCarApplyQueryAPIResponse() *AlitripBtripCityCarApplyQueryAPIResponse {
	return poolAlitripBtripCityCarApplyQueryAPIResponse.Get().(*AlitripBtripCityCarApplyQueryAPIResponse)
}

// ReleaseAlitripBtripCityCarApplyQueryAPIResponse 将 AlitripBtripCityCarApplyQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCityCarApplyQueryAPIResponse(v *AlitripBtripCityCarApplyQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripCityCarApplyQueryAPIResponse.Put(v)
}
