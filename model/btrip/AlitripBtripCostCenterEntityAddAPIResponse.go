package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterEntityAddAPIResponse 增加外部成本中心人员信息 API返回值
// alitrip.btrip.cost.center.entity.add
//
// 增加外部成本中心人员信息
type AlitripBtripCostCenterEntityAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterEntityAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterEntityAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterEntityAddAPIResponseModel).Reset()
}

// AlitripBtripCostCenterEntityAddAPIResponseModel is 增加外部成本中心人员信息 成功返回结果
type AlitripBtripCostCenterEntityAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_entity_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterEntityAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterEntityAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterEntityAddAPIResponse)
	},
}

// GetAlitripBtripCostCenterEntityAddAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterEntityAddAPIResponse
func GetAlitripBtripCostCenterEntityAddAPIResponse() *AlitripBtripCostCenterEntityAddAPIResponse {
	return poolAlitripBtripCostCenterEntityAddAPIResponse.Get().(*AlitripBtripCostCenterEntityAddAPIResponse)
}

// ReleaseAlitripBtripCostCenterEntityAddAPIResponse 将 AlitripBtripCostCenterEntityAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterEntityAddAPIResponse(v *AlitripBtripCostCenterEntityAddAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterEntityAddAPIResponse.Put(v)
}
