package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtipCostCenterQueryAPIResponse 查询外部成本中心 API返回值
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
type AlitripBtipCostCenterQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtipCostCenterQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtipCostCenterQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtipCostCenterQueryAPIResponseModel).Reset()
}

// AlitripBtipCostCenterQueryAPIResponseModel is 查询外部成本中心 成功返回结果
type AlitripBtipCostCenterQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btip_cost_center_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtipCostCenterQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtipCostCenterQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtipCostCenterQueryAPIResponse)
	},
}

// GetAlitripBtipCostCenterQueryAPIResponse 从 sync.Pool 获取 AlitripBtipCostCenterQueryAPIResponse
func GetAlitripBtipCostCenterQueryAPIResponse() *AlitripBtipCostCenterQueryAPIResponse {
	return poolAlitripBtipCostCenterQueryAPIResponse.Get().(*AlitripBtipCostCenterQueryAPIResponse)
}

// ReleaseAlitripBtipCostCenterQueryAPIResponse 将 AlitripBtipCostCenterQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtipCostCenterQueryAPIResponse(v *AlitripBtipCostCenterQueryAPIResponse) {
	v.Reset()
	poolAlitripBtipCostCenterQueryAPIResponse.Put(v)
}
