package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressModifyAPIResponse 卖家地址库修改 API返回值
// taobao.logistics.address.modify
//
// 卖家地址库修改
type TaobaoLogisticsAddressModifyAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAddressModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsAddressModifyAPIResponseModel).Reset()
}

// TaobaoLogisticsAddressModifyAPIResponseModel is 卖家地址库修改 成功返回结果
type TaobaoLogisticsAddressModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只返回修改时间modify_date
	AddressResult *AddressResult `json:"address_result,omitempty" xml:"address_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddressResult = nil
}

var poolTaobaoLogisticsAddressModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsAddressModifyAPIResponse)
	},
}

// GetTaobaoLogisticsAddressModifyAPIResponse 从 sync.Pool 获取 TaobaoLogisticsAddressModifyAPIResponse
func GetTaobaoLogisticsAddressModifyAPIResponse() *TaobaoLogisticsAddressModifyAPIResponse {
	return poolTaobaoLogisticsAddressModifyAPIResponse.Get().(*TaobaoLogisticsAddressModifyAPIResponse)
}

// ReleaseTaobaoLogisticsAddressModifyAPIResponse 将 TaobaoLogisticsAddressModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsAddressModifyAPIResponse(v *TaobaoLogisticsAddressModifyAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsAddressModifyAPIResponse.Put(v)
}
