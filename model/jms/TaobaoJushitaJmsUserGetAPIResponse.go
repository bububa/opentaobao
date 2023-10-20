package jms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserGetAPIResponse 查询某个用户是否同步消息 API返回值
// taobao.jushita.jms.user.get
//
// 查询某个用户是否同步消息，只支持单个查询
type TaobaoJushitaJmsUserGetAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsUserGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsUserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJushitaJmsUserGetAPIResponseModel).Reset()
}

// TaobaoJushitaJmsUserGetAPIResponseModel is 查询某个用户是否同步消息 成功返回结果
type TaobaoJushitaJmsUserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_user_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步的用户信息
	OnsUser *TmcUser `json:"ons_user,omitempty" xml:"ons_user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsUserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OnsUser = nil
}

var poolTaobaoJushitaJmsUserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJushitaJmsUserGetAPIResponse)
	},
}

// GetTaobaoJushitaJmsUserGetAPIResponse 从 sync.Pool 获取 TaobaoJushitaJmsUserGetAPIResponse
func GetTaobaoJushitaJmsUserGetAPIResponse() *TaobaoJushitaJmsUserGetAPIResponse {
	return poolTaobaoJushitaJmsUserGetAPIResponse.Get().(*TaobaoJushitaJmsUserGetAPIResponse)
}

// ReleaseTaobaoJushitaJmsUserGetAPIResponse 将 TaobaoJushitaJmsUserGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJushitaJmsUserGetAPIResponse(v *TaobaoJushitaJmsUserGetAPIResponse) {
	v.Reset()
	poolTaobaoJushitaJmsUserGetAPIResponse.Put(v)
}
