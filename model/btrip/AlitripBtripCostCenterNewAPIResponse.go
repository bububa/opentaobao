package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterNewAPIResponse 新建外部成本中心 API返回值
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
type AlitripBtripCostCenterNewAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterNewAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterNewAPIResponseModel).Reset()
}

// AlitripBtripCostCenterNewAPIResponseModel is 新建外部成本中心 成功返回结果
type AlitripBtripCostCenterNewAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterNewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterNewAPIResponse)
	},
}

// GetAlitripBtripCostCenterNewAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterNewAPIResponse
func GetAlitripBtripCostCenterNewAPIResponse() *AlitripBtripCostCenterNewAPIResponse {
	return poolAlitripBtripCostCenterNewAPIResponse.Get().(*AlitripBtripCostCenterNewAPIResponse)
}

// ReleaseAlitripBtripCostCenterNewAPIResponse 将 AlitripBtripCostCenterNewAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterNewAPIResponse(v *AlitripBtripCostCenterNewAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterNewAPIResponse.Put(v)
}
