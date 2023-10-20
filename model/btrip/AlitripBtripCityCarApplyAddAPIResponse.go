package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCityCarApplyAddAPIResponse 三方市内用车申请单同步 API返回值
// alitrip.btrip.city.car.apply.add
//
// 三方市内用车申请单同步
type AlitripBtripCityCarApplyAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripCityCarApplyAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCityCarApplyAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCityCarApplyAddAPIResponseModel).Reset()
}

// AlitripBtripCityCarApplyAddAPIResponseModel is 三方市内用车申请单同步 成功返回结果
type AlitripBtripCityCarApplyAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_city_car_apply_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 商旅内部审批单ID
	Module int64 `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCityCarApplyAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Module = 0
}

var poolAlitripBtripCityCarApplyAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCityCarApplyAddAPIResponse)
	},
}

// GetAlitripBtripCityCarApplyAddAPIResponse 从 sync.Pool 获取 AlitripBtripCityCarApplyAddAPIResponse
func GetAlitripBtripCityCarApplyAddAPIResponse() *AlitripBtripCityCarApplyAddAPIResponse {
	return poolAlitripBtripCityCarApplyAddAPIResponse.Get().(*AlitripBtripCityCarApplyAddAPIResponse)
}

// ReleaseAlitripBtripCityCarApplyAddAPIResponse 将 AlitripBtripCityCarApplyAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCityCarApplyAddAPIResponse(v *AlitripBtripCityCarApplyAddAPIResponse) {
	v.Reset()
	poolAlitripBtripCityCarApplyAddAPIResponse.Put(v)
}
