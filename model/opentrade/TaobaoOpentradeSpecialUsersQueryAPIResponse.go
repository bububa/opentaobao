package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialUsersQueryAPIResponse 专属下单标记信息查询 API返回值
// taobao.opentrade.special.users.query
//
// 专属下单标记信息查询
type TaobaoOpentradeSpecialUsersQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeSpecialUsersQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialUsersQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeSpecialUsersQueryAPIResponseModel).Reset()
}

// TaobaoOpentradeSpecialUsersQueryAPIResponseModel is 专属下单标记信息查询 成功返回结果
type TaobaoOpentradeSpecialUsersQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_special_users_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标记用户数据
	Results []MarkUserInfo `json:"results,omitempty" xml:"results>mark_user_info,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialUsersQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.TotalCount = 0
}

var poolTaobaoOpentradeSpecialUsersQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeSpecialUsersQueryAPIResponse)
	},
}

// GetTaobaoOpentradeSpecialUsersQueryAPIResponse 从 sync.Pool 获取 TaobaoOpentradeSpecialUsersQueryAPIResponse
func GetTaobaoOpentradeSpecialUsersQueryAPIResponse() *TaobaoOpentradeSpecialUsersQueryAPIResponse {
	return poolTaobaoOpentradeSpecialUsersQueryAPIResponse.Get().(*TaobaoOpentradeSpecialUsersQueryAPIResponse)
}

// ReleaseTaobaoOpentradeSpecialUsersQueryAPIResponse 将 TaobaoOpentradeSpecialUsersQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeSpecialUsersQueryAPIResponse(v *TaobaoOpentradeSpecialUsersQueryAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeSpecialUsersQueryAPIResponse.Put(v)
}
