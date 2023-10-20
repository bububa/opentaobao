package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse 上门取退可揽范围黑名单删除接口 API返回值
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
type TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponseModel is 上门取退可揽范围黑名单删除接口 成功返回结果
type TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_address_blacklist_tms_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	AddressBlacklistDeleteResponse *AddressBlacklistDeleteResponse `json:"address_blacklist_delete_response,omitempty" xml:"address_blacklist_delete_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddressBlacklistDeleteResponse = nil
}

var poolTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse)
	},
}

// GetTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse
func GetTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse() *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse {
	return poolTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse.Get().(*TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse)
}

// ReleaseTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse 将 TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse(v *TaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressAddressBlacklistTmsDeleteAPIResponse.Put(v)
}
