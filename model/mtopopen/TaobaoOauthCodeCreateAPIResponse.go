package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOauthCodeCreateAPIResponse 淘宝OauthCode颁发 API返回值
// taobao.oauth.code.create
//
// 手淘无线开放的oauthCode颁发接口
type TaobaoOauthCodeCreateAPIResponse struct {
	model.CommonResponse
	TaobaoOauthCodeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOauthCodeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOauthCodeCreateAPIResponseModel).Reset()
}

// TaobaoOauthCodeCreateAPIResponseModel is 淘宝OauthCode颁发 成功返回结果
type TaobaoOauthCodeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"oauth_code_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// mock out params
	Test int64 `json:"test,omitempty" xml:"test,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOauthCodeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Test = 0
}

var poolTaobaoOauthCodeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOauthCodeCreateAPIResponse)
	},
}

// GetTaobaoOauthCodeCreateAPIResponse 从 sync.Pool 获取 TaobaoOauthCodeCreateAPIResponse
func GetTaobaoOauthCodeCreateAPIResponse() *TaobaoOauthCodeCreateAPIResponse {
	return poolTaobaoOauthCodeCreateAPIResponse.Get().(*TaobaoOauthCodeCreateAPIResponse)
}

// ReleaseTaobaoOauthCodeCreateAPIResponse 将 TaobaoOauthCodeCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOauthCodeCreateAPIResponse(v *TaobaoOauthCodeCreateAPIResponse) {
	v.Reset()
	poolTaobaoOauthCodeCreateAPIResponse.Put(v)
}
