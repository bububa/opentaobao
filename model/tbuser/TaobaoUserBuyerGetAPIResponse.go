package tbuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserBuyerGetAPIResponse 查询买家信息API API返回值
// taobao.user.buyer.get
//
// 查询买家信息API，只能买家类应用调用。
type TaobaoUserBuyerGetAPIResponse struct {
	model.CommonResponse
	TaobaoUserBuyerGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUserBuyerGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUserBuyerGetAPIResponseModel).Reset()
}

// TaobaoUserBuyerGetAPIResponseModel is 查询买家信息API 成功返回结果
type TaobaoUserBuyerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_buyer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户
	User *User `json:"user,omitempty" xml:"user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUserBuyerGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.User = nil
}

var poolTaobaoUserBuyerGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUserBuyerGetAPIResponse)
	},
}

// GetTaobaoUserBuyerGetAPIResponse 从 sync.Pool 获取 TaobaoUserBuyerGetAPIResponse
func GetTaobaoUserBuyerGetAPIResponse() *TaobaoUserBuyerGetAPIResponse {
	return poolTaobaoUserBuyerGetAPIResponse.Get().(*TaobaoUserBuyerGetAPIResponse)
}

// ReleaseTaobaoUserBuyerGetAPIResponse 将 TaobaoUserBuyerGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUserBuyerGetAPIResponse(v *TaobaoUserBuyerGetAPIResponse) {
	v.Reset()
	poolTaobaoUserBuyerGetAPIResponse.Put(v)
}
