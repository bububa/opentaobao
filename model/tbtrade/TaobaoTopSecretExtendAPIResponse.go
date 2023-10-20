package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretExtendAPIResponse 虚拟号延期 API返回值
// taobao.top.secret.extend
//
// 虚拟号延期
type TaobaoTopSecretExtendAPIResponse struct {
	model.CommonResponse
	TaobaoTopSecretExtendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopSecretExtendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopSecretExtendAPIResponseModel).Reset()
}

// TaobaoTopSecretExtendAPIResponseModel is 虚拟号延期 成功返回结果
type TaobaoTopSecretExtendAPIResponseModel struct {
	XMLName xml.Name `xml:"top_secret_extend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 隐私号延期返回结果
	Result *SecretNo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopSecretExtendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTopSecretExtendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopSecretExtendAPIResponse)
	},
}

// GetTaobaoTopSecretExtendAPIResponse 从 sync.Pool 获取 TaobaoTopSecretExtendAPIResponse
func GetTaobaoTopSecretExtendAPIResponse() *TaobaoTopSecretExtendAPIResponse {
	return poolTaobaoTopSecretExtendAPIResponse.Get().(*TaobaoTopSecretExtendAPIResponse)
}

// ReleaseTaobaoTopSecretExtendAPIResponse 将 TaobaoTopSecretExtendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopSecretExtendAPIResponse(v *TaobaoTopSecretExtendAPIResponse) {
	v.Reset()
	poolTaobaoTopSecretExtendAPIResponse.Put(v)
}
