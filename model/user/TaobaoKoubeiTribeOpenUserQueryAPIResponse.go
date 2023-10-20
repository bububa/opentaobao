package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiTribeOpenUserQueryAPIResponse 获取用户openId API返回值
// taobao.koubei.tribe.open.user.query
//
// 口碑综合体通过手机号码获取加密后的用户openId
type TaobaoKoubeiTribeOpenUserQueryAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiTribeOpenUserQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiTribeOpenUserQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiTribeOpenUserQueryAPIResponseModel).Reset()
}

// TaobaoKoubeiTribeOpenUserQueryAPIResponseModel is 获取用户openId 成功返回结果
type TaobaoKoubeiTribeOpenUserQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_tribe_open_user_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiTribeOpenUserQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiTribeOpenUserQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiTribeOpenUserQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiTribeOpenUserQueryAPIResponse)
	},
}

// GetTaobaoKoubeiTribeOpenUserQueryAPIResponse 从 sync.Pool 获取 TaobaoKoubeiTribeOpenUserQueryAPIResponse
func GetTaobaoKoubeiTribeOpenUserQueryAPIResponse() *TaobaoKoubeiTribeOpenUserQueryAPIResponse {
	return poolTaobaoKoubeiTribeOpenUserQueryAPIResponse.Get().(*TaobaoKoubeiTribeOpenUserQueryAPIResponse)
}

// ReleaseTaobaoKoubeiTribeOpenUserQueryAPIResponse 将 TaobaoKoubeiTribeOpenUserQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiTribeOpenUserQueryAPIResponse(v *TaobaoKoubeiTribeOpenUserQueryAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiTribeOpenUserQueryAPIResponse.Put(v)
}
