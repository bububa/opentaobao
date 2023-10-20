package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse 上门取退用户黑名单同步 API返回值
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
type TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponseModel is 上门取退用户黑名单同步 成功返回结果
type TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_user_blacklist_tms_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	UserBlacklistResponse *UserBlacklistResponse `json:"user_blacklist_response,omitempty" xml:"user_blacklist_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UserBlacklistResponse = nil
}

var poolTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse
func GetTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse() *TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse {
	return poolTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse.Get().(*TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse 将 TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse(v *TaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressUserBlacklistTmsSyncAPIResponse.Put(v)
}
