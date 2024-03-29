package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterModifyAPIResponse 修改外部成本中心 API返回值
// alitrip.btrip.cost.center.modify
//
// 修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
type AlitripBtripCostCenterModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterModifyAPIResponseModel).Reset()
}

// AlitripBtripCostCenterModifyAPIResponseModel is 修改外部成本中心 成功返回结果
type AlitripBtripCostCenterModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterModifyAPIResponse)
	},
}

// GetAlitripBtripCostCenterModifyAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterModifyAPIResponse
func GetAlitripBtripCostCenterModifyAPIResponse() *AlitripBtripCostCenterModifyAPIResponse {
	return poolAlitripBtripCostCenterModifyAPIResponse.Get().(*AlitripBtripCostCenterModifyAPIResponse)
}

// ReleaseAlitripBtripCostCenterModifyAPIResponse 将 AlitripBtripCostCenterModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterModifyAPIResponse(v *AlitripBtripCostCenterModifyAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterModifyAPIResponse.Put(v)
}
