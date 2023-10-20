package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterDeleteAPIResponse 删除成本中心 API返回值
// alitrip.btrip.open.cost.center.delete
//
// 删除成本中心
type AlitripBtripOpenCostCenterDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterDeleteAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterDeleteAPIResponseModel is 删除成本中心 成功返回结果
type AlitripBtripOpenCostCenterDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterDeleteAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterDeleteAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterDeleteAPIResponse
func GetAlitripBtripOpenCostCenterDeleteAPIResponse() *AlitripBtripOpenCostCenterDeleteAPIResponse {
	return poolAlitripBtripOpenCostCenterDeleteAPIResponse.Get().(*AlitripBtripOpenCostCenterDeleteAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterDeleteAPIResponse 将 AlitripBtripOpenCostCenterDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterDeleteAPIResponse(v *AlitripBtripOpenCostCenterDeleteAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterDeleteAPIResponse.Put(v)
}
