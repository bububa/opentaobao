package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterTransferAPIResponse 商旅成本中心转换为外部成本中心 API返回值
// alitrip.btrip.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
type AlitripBtripCostCenterTransferAPIResponse struct {
	model.CommonResponse
	AlitripBtripCostCenterTransferAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterTransferAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCostCenterTransferAPIResponseModel).Reset()
}

// AlitripBtripCostCenterTransferAPIResponseModel is 商旅成本中心转换为外部成本中心 成功返回结果
type AlitripBtripCostCenterTransferAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_cost_center_transfer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BcmcResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCostCenterTransferAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCostCenterTransferAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCostCenterTransferAPIResponse)
	},
}

// GetAlitripBtripCostCenterTransferAPIResponse 从 sync.Pool 获取 AlitripBtripCostCenterTransferAPIResponse
func GetAlitripBtripCostCenterTransferAPIResponse() *AlitripBtripCostCenterTransferAPIResponse {
	return poolAlitripBtripCostCenterTransferAPIResponse.Get().(*AlitripBtripCostCenterTransferAPIResponse)
}

// ReleaseAlitripBtripCostCenterTransferAPIResponse 将 AlitripBtripCostCenterTransferAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCostCenterTransferAPIResponse(v *AlitripBtripCostCenterTransferAPIResponse) {
	v.Reset()
	poolAlitripBtripCostCenterTransferAPIResponse.Put(v)
}
