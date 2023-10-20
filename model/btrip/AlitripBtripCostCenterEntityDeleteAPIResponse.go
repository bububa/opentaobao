package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterEntityDeleteAPIResponse 删除外部成本中心人员信息 API返回值
// alitrip.btrip.cost.center.entity.delete
//
// 删除外部成本中心人员信息
type AlitripBtripCostCenterEntityDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterEntityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterEntityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterEntityDeleteAPIResponseModel).Reset()
}

// AlitripBtripCostCenterEntityDeleteAPIResponseModel is 删除外部成本中心人员信息 成功返回结果
type AlitripBtripCostCenterEntityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_entity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterEntityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterEntityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterEntityDeleteAPIResponse)
	},
}

// GetAlitripBtripCostCenterEntityDeleteAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterEntityDeleteAPIResponse
func GetAlitripBtripCostCenterEntityDeleteAPIResponse() *AlitripBtripCostCenterEntityDeleteAPIResponse {
	return poolAlitripBtripCostCenterEntityDeleteAPIResponse.Get().(*AlitripBtripCostCenterEntityDeleteAPIResponse)
}

// ReleaseAlitripBtripCostCenterEntityDeleteAPIResponse 将 AlitripBtripCostCenterEntityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterEntityDeleteAPIResponse(v *AlitripBtripCostCenterEntityDeleteAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterEntityDeleteAPIResponse.Put(v)
}
