package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTradeUsersGetAPIResponse 获取奇门用户列表 API返回值
// taobao.qimen.trade.users.get
//
// 获取已开通奇门订单服务的用户列表
type TaobaoQimenTradeUsersGetAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTradeUsersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenTradeUsersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenTradeUsersGetAPIResponseModel).Reset()
}

// TaobaoQimenTradeUsersGetAPIResponseModel is 获取奇门用户列表 成功返回结果
type TaobaoQimenTradeUsersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_trade_users_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// modal
	Users []QimenUser `json:"users,omitempty" xml:"users>qimen_user,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenTradeUsersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Users = m.Users[:0]
	m.TotalCount = 0
}

var poolTaobaoQimenTradeUsersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTradeUsersGetAPIResponse)
	},
}

// GetTaobaoQimenTradeUsersGetAPIResponse 从 sync.Pool 获取 TaobaoQimenTradeUsersGetAPIResponse
func GetTaobaoQimenTradeUsersGetAPIResponse() *TaobaoQimenTradeUsersGetAPIResponse {
	return poolTaobaoQimenTradeUsersGetAPIResponse.Get().(*TaobaoQimenTradeUsersGetAPIResponse)
}

// ReleaseTaobaoQimenTradeUsersGetAPIResponse 将 TaobaoQimenTradeUsersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenTradeUsersGetAPIResponse(v *TaobaoQimenTradeUsersGetAPIResponse) {
	v.Reset()
	poolTaobaoQimenTradeUsersGetAPIResponse.Put(v)
}
