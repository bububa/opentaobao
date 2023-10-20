package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressAddAPIResponse 卖家地址库新增接口 API返回值
// taobao.logistics.address.add
//
// 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
type TaobaoLogisticsAddressAddAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAddressAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsAddressAddAPIResponseModel).Reset()
}

// TaobaoLogisticsAddressAddAPIResponseModel is 卖家地址库新增接口 成功返回结果
type TaobaoLogisticsAddressAddAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只返回修改日期modify_date
	AddressResult *AddressResult `json:"address_result,omitempty" xml:"address_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddressResult = nil
}

var poolTaobaoLogisticsAddressAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsAddressAddAPIResponse)
	},
}

// GetTaobaoLogisticsAddressAddAPIResponse 从 sync.Pool 获取 TaobaoLogisticsAddressAddAPIResponse
func GetTaobaoLogisticsAddressAddAPIResponse() *TaobaoLogisticsAddressAddAPIResponse {
	return poolTaobaoLogisticsAddressAddAPIResponse.Get().(*TaobaoLogisticsAddressAddAPIResponse)
}

// ReleaseTaobaoLogisticsAddressAddAPIResponse 将 TaobaoLogisticsAddressAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsAddressAddAPIResponse(v *TaobaoLogisticsAddressAddAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsAddressAddAPIResponse.Put(v)
}
