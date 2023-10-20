package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressSearchAPIResponse 查询卖家地址库 API返回值
// taobao.logistics.address.search
//
// 通过此接口查询卖家地址库，
type TaobaoLogisticsAddressSearchAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAddressSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsAddressSearchAPIResponseModel).Reset()
}

// TaobaoLogisticsAddressSearchAPIResponseModel is 查询卖家地址库 成功返回结果
type TaobaoLogisticsAddressSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 一组地址库数据，
	Addresses []AddressResult `json:"addresses,omitempty" xml:"addresses>address_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Addresses = m.Addresses[:0]
}

var poolTaobaoLogisticsAddressSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsAddressSearchAPIResponse)
	},
}

// GetTaobaoLogisticsAddressSearchAPIResponse 从 sync.Pool 获取 TaobaoLogisticsAddressSearchAPIResponse
func GetTaobaoLogisticsAddressSearchAPIResponse() *TaobaoLogisticsAddressSearchAPIResponse {
	return poolTaobaoLogisticsAddressSearchAPIResponse.Get().(*TaobaoLogisticsAddressSearchAPIResponse)
}

// ReleaseTaobaoLogisticsAddressSearchAPIResponse 将 TaobaoLogisticsAddressSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsAddressSearchAPIResponse(v *TaobaoLogisticsAddressSearchAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsAddressSearchAPIResponse.Put(v)
}
