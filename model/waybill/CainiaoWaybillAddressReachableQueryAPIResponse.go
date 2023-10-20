package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillAddressReachableQueryAPIResponse 地址可达查询 API返回值
// cainiao.waybill.address.reachable.query
//
// 地址可达查询
type CainiaoWaybillAddressReachableQueryAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillAddressReachableQueryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillAddressReachableQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillAddressReachableQueryAPIResponseModel).Reset()
}

// CainiaoWaybillAddressReachableQueryAPIResponseModel is 地址可达查询 成功返回结果
type CainiaoWaybillAddressReachableQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_address_reachable_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillAddressReachableQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoWaybillAddressReachableQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillAddressReachableQueryAPIResponse)
	},
}

// GetCainiaoWaybillAddressReachableQueryAPIResponse 从 sync.Pool 获取 CainiaoWaybillAddressReachableQueryAPIResponse
func GetCainiaoWaybillAddressReachableQueryAPIResponse() *CainiaoWaybillAddressReachableQueryAPIResponse {
	return poolCainiaoWaybillAddressReachableQueryAPIResponse.Get().(*CainiaoWaybillAddressReachableQueryAPIResponse)
}

// ReleaseCainiaoWaybillAddressReachableQueryAPIResponse 将 CainiaoWaybillAddressReachableQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillAddressReachableQueryAPIResponse(v *CainiaoWaybillAddressReachableQueryAPIResponse) {
	v.Reset()
	poolCainiaoWaybillAddressReachableQueryAPIResponse.Put(v)
}
