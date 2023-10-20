package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialUsersMarkAPIResponse 专属下单可购买用户标记 API返回值
// taobao.opentrade.special.users.mark
//
// 对于专属下单的交易场景，根据openid标记用户可购买商品
type TaobaoOpentradeSpecialUsersMarkAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeSpecialUsersMarkAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialUsersMarkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeSpecialUsersMarkAPIResponseModel).Reset()
}

// TaobaoOpentradeSpecialUsersMarkAPIResponseModel is 专属下单可购买用户标记 成功返回结果
type TaobaoOpentradeSpecialUsersMarkAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_special_users_mark_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标记成功的用户数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialUsersMarkAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolTaobaoOpentradeSpecialUsersMarkAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeSpecialUsersMarkAPIResponse)
	},
}

// GetTaobaoOpentradeSpecialUsersMarkAPIResponse 从 sync.Pool 获取 TaobaoOpentradeSpecialUsersMarkAPIResponse
func GetTaobaoOpentradeSpecialUsersMarkAPIResponse() *TaobaoOpentradeSpecialUsersMarkAPIResponse {
	return poolTaobaoOpentradeSpecialUsersMarkAPIResponse.Get().(*TaobaoOpentradeSpecialUsersMarkAPIResponse)
}

// ReleaseTaobaoOpentradeSpecialUsersMarkAPIResponse 将 TaobaoOpentradeSpecialUsersMarkAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeSpecialUsersMarkAPIResponse(v *TaobaoOpentradeSpecialUsersMarkAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeSpecialUsersMarkAPIResponse.Put(v)
}
