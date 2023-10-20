package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenidGetAPIResponse 查询用户openId API返回值
// taobao.user.openid.get
//
// 查询用户openId
type TaobaoUserOpenidGetAPIResponse struct {
	model.CommonResponse
	TaobaoUserOpenidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUserOpenidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUserOpenidGetAPIResponseModel).Reset()
}

// TaobaoUserOpenidGetAPIResponseModel is 查询用户openId 成功返回结果
type TaobaoUserOpenidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_openid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对应账号的OpenUID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUserOpenidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenId = ""
}

var poolTaobaoUserOpenidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUserOpenidGetAPIResponse)
	},
}

// GetTaobaoUserOpenidGetAPIResponse 从 sync.Pool 获取 TaobaoUserOpenidGetAPIResponse
func GetTaobaoUserOpenidGetAPIResponse() *TaobaoUserOpenidGetAPIResponse {
	return poolTaobaoUserOpenidGetAPIResponse.Get().(*TaobaoUserOpenidGetAPIResponse)
}

// ReleaseTaobaoUserOpenidGetAPIResponse 将 TaobaoUserOpenidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUserOpenidGetAPIResponse(v *TaobaoUserOpenidGetAPIResponse) {
	v.Reset()
	poolTaobaoUserOpenidGetAPIResponse.Put(v)
}
