package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntitySetAPIResponse 设置成本中心人员信息 API返回值
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
type AlitripBtripOpenCostCenterEntitySetAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterEntitySetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterEntitySetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterEntitySetAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterEntitySetAPIResponseModel is 设置成本中心人员信息 成功返回结果
type AlitripBtripOpenCostCenterEntitySetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_entity_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Result *OpenCostCenterSetEntityRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterEntitySetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Result = nil
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterEntitySetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterEntitySetAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterEntitySetAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterEntitySetAPIResponse
func GetAlitripBtripOpenCostCenterEntitySetAPIResponse() *AlitripBtripOpenCostCenterEntitySetAPIResponse {
	return poolAlitripBtripOpenCostCenterEntitySetAPIResponse.Get().(*AlitripBtripOpenCostCenterEntitySetAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterEntitySetAPIResponse 将 AlitripBtripOpenCostCenterEntitySetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterEntitySetAPIResponse(v *AlitripBtripOpenCostCenterEntitySetAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterEntitySetAPIResponse.Put(v)
}
