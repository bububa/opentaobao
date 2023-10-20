package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterNewAPIResponse 新增成本中心 API返回值
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
type AlitripBtripOpenCostCenterNewAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterNewAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterNewAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterNewAPIResponseModel is 新增成本中心 成功返回结果
type AlitripBtripOpenCostCenterNewAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成本中心对象
	Module *OpenCostCenterSaveRs `json:"module,omitempty" xml:"module,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Module = nil
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterNewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterNewAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterNewAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterNewAPIResponse
func GetAlitripBtripOpenCostCenterNewAPIResponse() *AlitripBtripOpenCostCenterNewAPIResponse {
	return poolAlitripBtripOpenCostCenterNewAPIResponse.Get().(*AlitripBtripOpenCostCenterNewAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterNewAPIResponse 将 AlitripBtripOpenCostCenterNewAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterNewAPIResponse(v *AlitripBtripOpenCostCenterNewAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterNewAPIResponse.Put(v)
}
