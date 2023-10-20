package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressuserblacklisttmssyncAPIResponse 上门取退用户黑名单同步 API返回值
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
type TaobaologisticsexpressuserblacklisttmssyncAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpressuserblacklisttmssyncAPIResponseModel
}

// TaobaologisticsexpressuserblacklisttmssyncAPIResponseModel is 上门取退用户黑名单同步 成功返回结果
type TaobaologisticsexpressuserblacklisttmssyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_user_blacklist_tms_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	UserBlacklistResponse *UserBlacklistResponse `json:"user_blacklist_response,omitempty" xml:"user_blacklist_response,omitempty"`
}
