package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCityCarApplyApproveAPIResponse 三方市内用车申请单审批 API返回值
// alitrip.btrip.city.car.apply.approve
//
// 三方市内用车申请单审批
type AlitripBtripCityCarApplyApproveAPIResponse struct {
	model.CommonResponse
	AlitripBtripCityCarApplyApproveAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCityCarApplyApproveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCityCarApplyApproveAPIResponseModel).Reset()
}

// AlitripBtripCityCarApplyApproveAPIResponseModel is 三方市内用车申请单审批 成功返回结果
type AlitripBtripCityCarApplyApproveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_city_car_apply_approve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 审批是否成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCityCarApplyApproveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Module = false
}

var poolAlitripBtripCityCarApplyApproveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCityCarApplyApproveAPIResponse)
	},
}

// GetAlitripBtripCityCarApplyApproveAPIResponse 从 sync.Pool 获取 AlitripBtripCityCarApplyApproveAPIResponse
func GetAlitripBtripCityCarApplyApproveAPIResponse() *AlitripBtripCityCarApplyApproveAPIResponse {
	return poolAlitripBtripCityCarApplyApproveAPIResponse.Get().(*AlitripBtripCityCarApplyApproveAPIResponse)
}

// ReleaseAlitripBtripCityCarApplyApproveAPIResponse 将 AlitripBtripCityCarApplyApproveAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCityCarApplyApproveAPIResponse(v *AlitripBtripCityCarApplyApproveAPIResponse) {
	v.Reset()
	poolAlitripBtripCityCarApplyApproveAPIResponse.Put(v)
}
