package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJdpUserDeleteAPIResponse 删除数据推送用户 API返回值
// taobao.jushita.jdp.user.delete
//
// 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJdpUserDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJushitaJdpUserDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJushitaJdpUserDeleteAPIResponseModel).Reset()
}

// TaobaoJushitaJdpUserDeleteAPIResponseModel is 删除数据推送用户 成功返回结果
type TaobaoJushitaJdpUserDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jdp_user_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJushitaJdpUserDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoJushitaJdpUserDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJushitaJdpUserDeleteAPIResponse)
	},
}

// GetTaobaoJushitaJdpUserDeleteAPIResponse 从 sync.Pool 获取 TaobaoJushitaJdpUserDeleteAPIResponse
func GetTaobaoJushitaJdpUserDeleteAPIResponse() *TaobaoJushitaJdpUserDeleteAPIResponse {
	return poolTaobaoJushitaJdpUserDeleteAPIResponse.Get().(*TaobaoJushitaJdpUserDeleteAPIResponse)
}

// ReleaseTaobaoJushitaJdpUserDeleteAPIResponse 将 TaobaoJushitaJdpUserDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJushitaJdpUserDeleteAPIResponse(v *TaobaoJushitaJdpUserDeleteAPIResponse) {
	v.Reset()
	poolTaobaoJushitaJdpUserDeleteAPIResponse.Put(v)
}
