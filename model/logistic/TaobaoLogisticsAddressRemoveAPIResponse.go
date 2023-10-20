package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressRemoveAPIResponse 删除卖家地址库 API返回值
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
type TaobaoLogisticsAddressRemoveAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAddressRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsAddressRemoveAPIResponseModel).Reset()
}

// TaobaoLogisticsAddressRemoveAPIResponseModel is 删除卖家地址库 成功返回结果
type TaobaoLogisticsAddressRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只返回修改日期modify_date
	AddressResult *AddressResult `json:"address_result,omitempty" xml:"address_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsAddressRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddressResult = nil
}

var poolTaobaoLogisticsAddressRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsAddressRemoveAPIResponse)
	},
}

// GetTaobaoLogisticsAddressRemoveAPIResponse 从 sync.Pool 获取 TaobaoLogisticsAddressRemoveAPIResponse
func GetTaobaoLogisticsAddressRemoveAPIResponse() *TaobaoLogisticsAddressRemoveAPIResponse {
	return poolTaobaoLogisticsAddressRemoveAPIResponse.Get().(*TaobaoLogisticsAddressRemoveAPIResponse)
}

// ReleaseTaobaoLogisticsAddressRemoveAPIResponse 将 TaobaoLogisticsAddressRemoveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsAddressRemoveAPIResponse(v *TaobaoLogisticsAddressRemoveAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsAddressRemoveAPIResponse.Put(v)
}
