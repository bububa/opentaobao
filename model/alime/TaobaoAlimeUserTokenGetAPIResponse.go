package alime

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlimeUserTokenGetAPIResponse 获取用户免登录令牌 API返回值
// taobao.alime.user.token.get
//
// 根据第三账号信息获取用户的免登录令牌
type TaobaoAlimeUserTokenGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlimeUserTokenGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlimeUserTokenGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlimeUserTokenGetAPIResponseModel).Reset()
}

// TaobaoAlimeUserTokenGetAPIResponseModel is 获取用户免登录令牌 成功返回结果
type TaobaoAlimeUserTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alime_user_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 响应消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应编码(由于&#34;code&#34;为top保留字，用code0以示区分，文档中均以code说明)，code == 0为成功，其它为失败
	Code0 int64 `json:"code0,omitempty" xml:"code0,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlimeUserTokenGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.Code0 = 0
}

var poolTaobaoAlimeUserTokenGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlimeUserTokenGetAPIResponse)
	},
}

// GetTaobaoAlimeUserTokenGetAPIResponse 从 sync.Pool 获取 TaobaoAlimeUserTokenGetAPIResponse
func GetTaobaoAlimeUserTokenGetAPIResponse() *TaobaoAlimeUserTokenGetAPIResponse {
	return poolTaobaoAlimeUserTokenGetAPIResponse.Get().(*TaobaoAlimeUserTokenGetAPIResponse)
}

// ReleaseTaobaoAlimeUserTokenGetAPIResponse 将 TaobaoAlimeUserTokenGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlimeUserTokenGetAPIResponse(v *TaobaoAlimeUserTokenGetAPIResponse) {
	v.Reset()
	poolTaobaoAlimeUserTokenGetAPIResponse.Put(v)
}
