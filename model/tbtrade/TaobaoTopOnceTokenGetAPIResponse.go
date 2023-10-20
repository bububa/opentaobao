package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOnceTokenGetAPIResponse 网关一次性token获取 API返回值
// taobao.top.once.token.get
//
// 网关一次性token获取，对接文档:
type TaobaoTopOnceTokenGetAPIResponse struct {
	model.CommonResponse
	TaobaoTopOnceTokenGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopOnceTokenGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopOnceTokenGetAPIResponseModel).Reset()
}

// TaobaoTopOnceTokenGetAPIResponseModel is 网关一次性token获取 成功返回结果
type TaobaoTopOnceTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"top_once_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 失败详情
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopOnceTokenGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Token = ""
	m.ResultMsg = ""
}

var poolTaobaoTopOnceTokenGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopOnceTokenGetAPIResponse)
	},
}

// GetTaobaoTopOnceTokenGetAPIResponse 从 sync.Pool 获取 TaobaoTopOnceTokenGetAPIResponse
func GetTaobaoTopOnceTokenGetAPIResponse() *TaobaoTopOnceTokenGetAPIResponse {
	return poolTaobaoTopOnceTokenGetAPIResponse.Get().(*TaobaoTopOnceTokenGetAPIResponse)
}

// ReleaseTaobaoTopOnceTokenGetAPIResponse 将 TaobaoTopOnceTokenGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopOnceTokenGetAPIResponse(v *TaobaoTopOnceTokenGetAPIResponse) {
	v.Reset()
	poolTaobaoTopOnceTokenGetAPIResponse.Put(v)
}
