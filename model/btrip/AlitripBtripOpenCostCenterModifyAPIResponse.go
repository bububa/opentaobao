package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterModifyAPIResponse 修改成本中心 API返回值
// alitrip.btrip.open.cost.center.modify
//
// 修改成本中心
type AlitripBtripOpenCostCenterModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterModifyAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterModifyAPIResponseModel is 修改成本中心 成功返回结果
type AlitripBtripOpenCostCenterModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterModifyAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterModifyAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterModifyAPIResponse
func GetAlitripBtripOpenCostCenterModifyAPIResponse() *AlitripBtripOpenCostCenterModifyAPIResponse {
	return poolAlitripBtripOpenCostCenterModifyAPIResponse.Get().(*AlitripBtripOpenCostCenterModifyAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterModifyAPIResponse 将 AlitripBtripOpenCostCenterModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterModifyAPIResponse(v *AlitripBtripOpenCostCenterModifyAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterModifyAPIResponse.Put(v)
}
