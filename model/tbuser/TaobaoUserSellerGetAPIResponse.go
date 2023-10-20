package tbuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserSellerGetAPIResponse 查询卖家用户信息 API返回值
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaoUserSellerGetAPIResponse struct {
	model.CommonResponse
	TaobaoUserSellerGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUserSellerGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUserSellerGetAPIResponseModel).Reset()
}

// TaobaoUserSellerGetAPIResponseModel is 查询卖家用户信息 成功返回结果
type TaobaoUserSellerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_seller_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户
	User *User `json:"user,omitempty" xml:"user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUserSellerGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.User = nil
}

var poolTaobaoUserSellerGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUserSellerGetAPIResponse)
	},
}

// GetTaobaoUserSellerGetAPIResponse 从 sync.Pool 获取 TaobaoUserSellerGetAPIResponse
func GetTaobaoUserSellerGetAPIResponse() *TaobaoUserSellerGetAPIResponse {
	return poolTaobaoUserSellerGetAPIResponse.Get().(*TaobaoUserSellerGetAPIResponse)
}

// ReleaseTaobaoUserSellerGetAPIResponse 将 TaobaoUserSellerGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUserSellerGetAPIResponse(v *TaobaoUserSellerGetAPIResponse) {
	v.Reset()
	poolTaobaoUserSellerGetAPIResponse.Put(v)
}
