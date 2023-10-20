package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterTransferAPIResponse 商旅成本中心转换为外部成本中心 API返回值
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
type AlitripBtripOpenCostCenterTransferAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenCostCenterTransferAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterTransferAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenCostCenterTransferAPIResponseModel).Reset()
}

// AlitripBtripOpenCostCenterTransferAPIResponseModel is 商旅成本中心转换为外部成本中心 成功返回结果
type AlitripBtripOpenCostCenterTransferAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_cost_center_transfer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenCostCenterTransferAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
}

var poolAlitripBtripOpenCostCenterTransferAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenCostCenterTransferAPIResponse)
	},
}

// GetAlitripBtripOpenCostCenterTransferAPIResponse 从 sync.Pool 获取 AlitripBtripOpenCostCenterTransferAPIResponse
func GetAlitripBtripOpenCostCenterTransferAPIResponse() *AlitripBtripOpenCostCenterTransferAPIResponse {
	return poolAlitripBtripOpenCostCenterTransferAPIResponse.Get().(*AlitripBtripOpenCostCenterTransferAPIResponse)
}

// ReleaseAlitripBtripOpenCostCenterTransferAPIResponse 将 AlitripBtripOpenCostCenterTransferAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenCostCenterTransferAPIResponse(v *AlitripBtripOpenCostCenterTransferAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenCostCenterTransferAPIResponse.Put(v)
}
