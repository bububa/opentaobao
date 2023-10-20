package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterGetAPIResponse 获取用户费用归属 API返回值
// alitrip.btrip.cost.center.get
//
// 获取差旅申请用户的费用归属列表
type AlitripBtripCostCenterGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterGetAPIResponseModel).Reset()
}

// AlitripBtripCostCenterGetAPIResponseModel is 获取用户费用归属 成功返回结果
type AlitripBtripCostCenterGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterGetAPIResponse)
	},
}

// GetAlitripBtripCostCenterGetAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterGetAPIResponse
func GetAlitripBtripCostCenterGetAPIResponse() *AlitripBtripCostCenterGetAPIResponse {
	return poolAlitripBtripCostCenterGetAPIResponse.Get().(*AlitripBtripCostCenterGetAPIResponse)
}

// ReleaseAlitripBtripCostCenterGetAPIResponse 将 AlitripBtripCostCenterGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterGetAPIResponse(v *AlitripBtripCostCenterGetAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterGetAPIResponse.Put(v)
}
