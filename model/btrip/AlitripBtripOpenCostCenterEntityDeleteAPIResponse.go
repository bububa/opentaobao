package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntityDeleteAPIResponse 删除成本中心人员信息 API返回值
// alitrip.btrip.open.cost.center.entity.delete
//
// 删除成本中心人员信息
type AlitripBtripOpenCostCenterEntityDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterEntityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterEntityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterEntityDeleteAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterEntityDeleteAPIResponseModel is 删除成本中心人员信息 成功返回结果
type AlitripBtripOpenCostCenterEntityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_entity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Result *OpenCostCenterDeleteEntityRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterEntityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Result = nil
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterEntityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterEntityDeleteAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterEntityDeleteAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterEntityDeleteAPIResponse
func GetAlitripBtripOpenCostCenterEntityDeleteAPIResponse() *AlitripBtripOpenCostCenterEntityDeleteAPIResponse {
	return poolAlitripBtripOpenCostCenterEntityDeleteAPIResponse.Get().(*AlitripBtripOpenCostCenterEntityDeleteAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterEntityDeleteAPIResponse 将 AlitripBtripOpenCostCenterEntityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterEntityDeleteAPIResponse(v *AlitripBtripOpenCostCenterEntityDeleteAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterEntityDeleteAPIResponse.Put(v)
}
