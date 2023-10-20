package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntityAddAPIResponse 增加成本中心人员信息 API返回值
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
type AlitripBtripOpenCostCenterEntityAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterEntityAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterEntityAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterEntityAddAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterEntityAddAPIResponseModel is 增加成本中心人员信息 成功返回结果
type AlitripBtripOpenCostCenterEntityAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_entity_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Result *OpenCostCenterAddEntityRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterEntityAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Result = nil
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterEntityAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterEntityAddAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterEntityAddAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterEntityAddAPIResponse
func GetAlitripBtripOpenCostCenterEntityAddAPIResponse() *AlitripBtripOpenCostCenterEntityAddAPIResponse {
	return poolAlitripBtripOpenCostCenterEntityAddAPIResponse.Get().(*AlitripBtripOpenCostCenterEntityAddAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterEntityAddAPIResponse 将 AlitripBtripOpenCostCenterEntityAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterEntityAddAPIResponse(v *AlitripBtripOpenCostCenterEntityAddAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterEntityAddAPIResponse.Put(v)
}
