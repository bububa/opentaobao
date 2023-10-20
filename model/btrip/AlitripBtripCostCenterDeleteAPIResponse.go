package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterDeleteAPIResponse 删除外部成本中心 API返回值
// alitrip.btrip.cost.center.delete
//
// 删除外部成本中心
type AlitripBtripCostCenterDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterDeleteAPIResponseModel).Reset()
}

// AlitripBtripCostCenterDeleteAPIResponseModel is 删除外部成本中心 成功返回结果
type AlitripBtripCostCenterDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterDeleteAPIResponse)
	},
}

// GetAlitripBtripCostCenterDeleteAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterDeleteAPIResponse
func GetAlitripBtripCostCenterDeleteAPIResponse() *AlitripBtripCostCenterDeleteAPIResponse {
	return poolAlitripBtripCostCenterDeleteAPIResponse.Get().(*AlitripBtripCostCenterDeleteAPIResponse)
}

// ReleaseAlitripBtripCostCenterDeleteAPIResponse 将 AlitripBtripCostCenterDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterDeleteAPIResponse(v *AlitripBtripCostCenterDeleteAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterDeleteAPIResponse.Put(v)
}
