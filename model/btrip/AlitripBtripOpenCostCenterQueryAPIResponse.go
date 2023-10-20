package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterQueryAPIResponse 查询成本中心 API返回值
// alitrip.btrip.open.cost.center.query
//
// 查询成本中心
type AlitripBtripOpenCostCenterQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterQueryAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterQueryAPIResponseModel is 查询成本中心 成功返回结果
type AlitripBtripOpenCostCenterQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成本中心列表
	CostCenterList []OpenCostCenterQueryRs `json:"cost_center_list,omitempty" xml:"cost_center_list>open_cost_center_query_rs,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CostCenterList = m.CostCenterList[:0]
	m.ResultMsg = ""
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterQueryAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterQueryAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterQueryAPIResponse
func GetAlitripBtripOpenCostCenterQueryAPIResponse() *AlitripBtripOpenCostCenterQueryAPIResponse {
	return poolAlitripBtripOpenCostCenterQueryAPIResponse.Get().(*AlitripBtripOpenCostCenterQueryAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterQueryAPIResponse 将 AlitripBtripOpenCostCenterQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterQueryAPIResponse(v *AlitripBtripOpenCostCenterQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterQueryAPIResponse.Put(v)
}
