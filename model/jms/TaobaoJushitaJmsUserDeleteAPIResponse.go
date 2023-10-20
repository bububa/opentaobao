package jms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserDeleteAPIResponse 删除ONS消息同步用户 API返回值
// taobao.jushita.jms.user.delete
//
// 删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
type TaobaoJushitaJmsUserDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsUserDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsUserDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJushitaJmsUserDeleteAPIResponseModel).Reset()
}

// TaobaoJushitaJmsUserDeleteAPIResponseModel is 删除ONS消息同步用户 成功返回结果
type TaobaoJushitaJmsUserDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_user_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsUserDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoJushitaJmsUserDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJushitaJmsUserDeleteAPIResponse)
	},
}

// GetTaobaoJushitaJmsUserDeleteAPIResponse 从 sync.Pool 获取 TaobaoJushitaJmsUserDeleteAPIResponse
func GetTaobaoJushitaJmsUserDeleteAPIResponse() *TaobaoJushitaJmsUserDeleteAPIResponse {
	return poolTaobaoJushitaJmsUserDeleteAPIResponse.Get().(*TaobaoJushitaJmsUserDeleteAPIResponse)
}

// ReleaseTaobaoJushitaJmsUserDeleteAPIResponse 将 TaobaoJushitaJmsUserDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJushitaJmsUserDeleteAPIResponse(v *TaobaoJushitaJmsUserDeleteAPIResponse) {
	v.Reset()
	poolTaobaoJushitaJmsUserDeleteAPIResponse.Put(v)
}
