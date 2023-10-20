package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponse 上门取退可揽范围黑名单同步/更新 API返回值
// taobao.logistics.express.address.blacklist.tms.async
//
// 上门取退可揽范围黑名单同步/更新
type TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponseModel
}

// TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponseModel is 上门取退可揽范围黑名单同步/更新 成功返回结果
type TaobaoLogisticsExpressAddressBlacklistTmsAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_address_blacklist_tms_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	AddressBlacklistResponse *AddressBlacklistResponse `json:"address_blacklist_response,omitempty" xml:"address_blacklist_response,omitempty"`
}
