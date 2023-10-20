package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoCarrierCapacityQueryAPIResponse 按照门店查询骑手运力状态查询 API返回值
// alibaba.ele.fengniao.carrier.capacity.query
//
// 提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
type AlibabaEleFengniaoCarrierCapacityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoCarrierCapacityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel).Reset()
}

// AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel is 按照门店查询骑手运力状态查询 成功返回结果
type AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_carrier_capacity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Results []Capacities `json:"results,omitempty" xml:"results>capacities,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoCarrierCapacityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolAlibabaEleFengniaoCarrierCapacityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoCarrierCapacityQueryAPIResponse)
	},
}

// GetAlibabaEleFengniaoCarrierCapacityQueryAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoCarrierCapacityQueryAPIResponse
func GetAlibabaEleFengniaoCarrierCapacityQueryAPIResponse() *AlibabaEleFengniaoCarrierCapacityQueryAPIResponse {
	return poolAlibabaEleFengniaoCarrierCapacityQueryAPIResponse.Get().(*AlibabaEleFengniaoCarrierCapacityQueryAPIResponse)
}

// ReleaseAlibabaEleFengniaoCarrierCapacityQueryAPIResponse 将 AlibabaEleFengniaoCarrierCapacityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoCarrierCapacityQueryAPIResponse(v *AlibabaEleFengniaoCarrierCapacityQueryAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoCarrierCapacityQueryAPIResponse.Put(v)
}
