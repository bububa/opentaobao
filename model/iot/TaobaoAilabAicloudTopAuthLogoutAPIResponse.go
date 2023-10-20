package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopAuthLogoutAPIResponse 登出 API返回值
// taobao.ailab.aicloud.top.auth.logout
//
// 登出
type TaobaoAilabAicloudTopAuthLogoutAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopAuthLogoutAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopAuthLogoutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopAuthLogoutAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopAuthLogoutAPIResponseModel is 登出 成功返回结果
type TaobaoAilabAicloudTopAuthLogoutAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_auth_logout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo错误码信息，成功返回null
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopAuthLogoutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopAuthLogoutAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopAuthLogoutAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopAuthLogoutAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopAuthLogoutAPIResponse
func GetTaobaoAilabAicloudTopAuthLogoutAPIResponse() *TaobaoAilabAicloudTopAuthLogoutAPIResponse {
	return poolTaobaoAilabAicloudTopAuthLogoutAPIResponse.Get().(*TaobaoAilabAicloudTopAuthLogoutAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopAuthLogoutAPIResponse 将 TaobaoAilabAicloudTopAuthLogoutAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopAuthLogoutAPIResponse(v *TaobaoAilabAicloudTopAuthLogoutAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopAuthLogoutAPIResponse.Put(v)
}
