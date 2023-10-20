package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaLoginAuthsignGetAPIResponse 获取登陆权限签名 API返回值
// taobao.simba.login.authsign.get
//
// 获取登陆权限签名
type TaobaoSimbaLoginAuthsignGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaLoginAuthsignGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaLoginAuthsignGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaLoginAuthsignGetAPIResponseModel).Reset()
}

// TaobaoSimbaLoginAuthsignGetAPIResponseModel is 获取登陆权限签名 成功返回结果
type TaobaoSimbaLoginAuthsignGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_login_authsign_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 登陆签名
	SubwayToken string `json:"subway_token,omitempty" xml:"subway_token,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaLoginAuthsignGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubwayToken = ""
}

var poolTaobaoSimbaLoginAuthsignGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaLoginAuthsignGetAPIResponse)
	},
}

// GetTaobaoSimbaLoginAuthsignGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaLoginAuthsignGetAPIResponse
func GetTaobaoSimbaLoginAuthsignGetAPIResponse() *TaobaoSimbaLoginAuthsignGetAPIResponse {
	return poolTaobaoSimbaLoginAuthsignGetAPIResponse.Get().(*TaobaoSimbaLoginAuthsignGetAPIResponse)
}

// ReleaseTaobaoSimbaLoginAuthsignGetAPIResponse 将 TaobaoSimbaLoginAuthsignGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaLoginAuthsignGetAPIResponse(v *TaobaoSimbaLoginAuthsignGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaLoginAuthsignGetAPIResponse.Put(v)
}
