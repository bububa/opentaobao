package alime

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlimeUserTokenAdvanceGetAPIResponse 获取用户免登录令牌v2 API返回值
// taobao.alime.user.token.advance.get
//
// 根据第三账号信息获取用户的免登录令牌
type TaobaoAlimeUserTokenAdvanceGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlimeUserTokenAdvanceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlimeUserTokenAdvanceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlimeUserTokenAdvanceGetAPIResponseModel).Reset()
}

// TaobaoAlimeUserTokenAdvanceGetAPIResponseModel is 获取用户免登录令牌v2 成功返回结果
type TaobaoAlimeUserTokenAdvanceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alime_user_token_advance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 响应编码(由于&#34;code&#34;为top保留字，用code0以示区分，文档中均以code说明)，code == 0为成功，其它为失败
	Code0 int64 `json:"code0,omitempty" xml:"code0,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlimeUserTokenAdvanceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Data = ""
	m.Code0 = 0
}

var poolTaobaoAlimeUserTokenAdvanceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlimeUserTokenAdvanceGetAPIResponse)
	},
}

// GetTaobaoAlimeUserTokenAdvanceGetAPIResponse 从 sync.Pool 获取 TaobaoAlimeUserTokenAdvanceGetAPIResponse
func GetTaobaoAlimeUserTokenAdvanceGetAPIResponse() *TaobaoAlimeUserTokenAdvanceGetAPIResponse {
	return poolTaobaoAlimeUserTokenAdvanceGetAPIResponse.Get().(*TaobaoAlimeUserTokenAdvanceGetAPIResponse)
}

// ReleaseTaobaoAlimeUserTokenAdvanceGetAPIResponse 将 TaobaoAlimeUserTokenAdvanceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlimeUserTokenAdvanceGetAPIResponse(v *TaobaoAlimeUserTokenAdvanceGetAPIResponse) {
	v.Reset()
	poolTaobaoAlimeUserTokenAdvanceGetAPIResponse.Put(v)
}
