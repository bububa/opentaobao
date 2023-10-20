package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterEntitySetAPIResponse 设置外部成本中心人员信息 API返回值
// alitrip.btrip.cost.center.entity.set
//
// 设置外部成本中心人员信息
type AlitripBtripCostCenterEntitySetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterEntitySetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterEntitySetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterEntitySetAPIResponseModel).Reset()
}

// AlitripBtripCostCenterEntitySetAPIResponseModel is 设置外部成本中心人员信息 成功返回结果
type AlitripBtripCostCenterEntitySetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_entity_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterEntitySetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterEntitySetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterEntitySetAPIResponse)
	},
}

// GetAlitripBtripCostCenterEntitySetAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterEntitySetAPIResponse
func GetAlitripBtripCostCenterEntitySetAPIResponse() *AlitripBtripCostCenterEntitySetAPIResponse {
	return poolAlitripBtripCostCenterEntitySetAPIResponse.Get().(*AlitripBtripCostCenterEntitySetAPIResponse)
}

// ReleaseAlitripBtripCostCenterEntitySetAPIResponse 将 AlitripBtripCostCenterEntitySetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterEntitySetAPIResponse(v *AlitripBtripCostCenterEntitySetAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterEntitySetAPIResponse.Put(v)
}
