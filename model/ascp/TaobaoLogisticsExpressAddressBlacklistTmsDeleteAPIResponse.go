package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponse 上门取退可揽范围黑名单删除接口 API返回值
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
type TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponseModel
}

// TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponseModel is 上门取退可揽范围黑名单删除接口 成功返回结果
type TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_address_blacklist_tms_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	AddressBlacklistDeleteResponse *AddressBlacklistDeleteResponse `json:"address_blacklist_delete_response,omitempty" xml:"address_blacklist_delete_response,omitempty"`
}
