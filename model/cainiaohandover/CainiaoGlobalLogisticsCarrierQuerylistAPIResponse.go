package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalLogisticsCarrierQuerylistAPIResponse 实际承运商查询 API返回值
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
type CainiaoGlobalLogisticsCarrierQuerylistAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalLogisticsCarrierQuerylistAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalLogisticsCarrierQuerylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalLogisticsCarrierQuerylistAPIResponseModel).Reset()
}

// CainiaoGlobalLogisticsCarrierQuerylistAPIResponseModel is 实际承运商查询 成功返回结果
type CainiaoGlobalLogisticsCarrierQuerylistAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_logistics_carrier_querylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *DubboResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalLogisticsCarrierQuerylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalLogisticsCarrierQuerylistAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalLogisticsCarrierQuerylistAPIResponse)
	},
}

// GetCainiaoGlobalLogisticsCarrierQuerylistAPIResponse 从 sync.Pool 获取 CainiaoGlobalLogisticsCarrierQuerylistAPIResponse
func GetCainiaoGlobalLogisticsCarrierQuerylistAPIResponse() *CainiaoGlobalLogisticsCarrierQuerylistAPIResponse {
	return poolCainiaoGlobalLogisticsCarrierQuerylistAPIResponse.Get().(*CainiaoGlobalLogisticsCarrierQuerylistAPIResponse)
}

// ReleaseCainiaoGlobalLogisticsCarrierQuerylistAPIResponse 将 CainiaoGlobalLogisticsCarrierQuerylistAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalLogisticsCarrierQuerylistAPIResponse(v *CainiaoGlobalLogisticsCarrierQuerylistAPIResponse) {
	v.Reset()
	poolCainiaoGlobalLogisticsCarrierQuerylistAPIResponse.Put(v)
}
