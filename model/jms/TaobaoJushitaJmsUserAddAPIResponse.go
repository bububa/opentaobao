package jms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserAddAPIResponse 添加ONS消息同步用户 API返回值
// taobao.jushita.jms.user.add
//
// 添加ONS消息同步用户
type TaobaoJushitaJmsUserAddAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsUserAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsUserAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJushitaJmsUserAddAPIResponseModel).Reset()
}

// TaobaoJushitaJmsUserAddAPIResponseModel is 添加ONS消息同步用户 成功返回结果
type TaobaoJushitaJmsUserAddAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_user_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功，如果失败请看错误信息
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsUserAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoJushitaJmsUserAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJushitaJmsUserAddAPIResponse)
	},
}

// GetTaobaoJushitaJmsUserAddAPIResponse 从 sync.Pool 获取 TaobaoJushitaJmsUserAddAPIResponse
func GetTaobaoJushitaJmsUserAddAPIResponse() *TaobaoJushitaJmsUserAddAPIResponse {
	return poolTaobaoJushitaJmsUserAddAPIResponse.Get().(*TaobaoJushitaJmsUserAddAPIResponse)
}

// ReleaseTaobaoJushitaJmsUserAddAPIResponse 将 TaobaoJushitaJmsUserAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJushitaJmsUserAddAPIResponse(v *TaobaoJushitaJmsUserAddAPIResponse) {
	v.Reset()
	poolTaobaoJushitaJmsUserAddAPIResponse.Put(v)
}
