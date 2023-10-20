package jstsecret

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSecretGetAPIResponse 获取订单消费者的隐私号码 API返回值
// taobao.jst.secret.get
//
// 根据订单号获取消费者的隐私号
type TaobaoJstSecretGetAPIResponse struct {
	model.CommonResponse
	TaobaoJstSecretGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSecretGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSecretGetAPIResponseModel).Reset()
}

// TaobaoJstSecretGetAPIResponseModel is 获取订单消费者的隐私号码 成功返回结果
type TaobaoJstSecretGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_secret_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Secret *SecretNoBindResponse `json:"secret,omitempty" xml:"secret,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstSecretGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Secret = nil
}

var poolTaobaoJstSecretGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSecretGetAPIResponse)
	},
}

// GetTaobaoJstSecretGetAPIResponse 从 sync.Pool 获取 TaobaoJstSecretGetAPIResponse
func GetTaobaoJstSecretGetAPIResponse() *TaobaoJstSecretGetAPIResponse {
	return poolTaobaoJstSecretGetAPIResponse.Get().(*TaobaoJstSecretGetAPIResponse)
}

// ReleaseTaobaoJstSecretGetAPIResponse 将 TaobaoJstSecretGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSecretGetAPIResponse(v *TaobaoJstSecretGetAPIResponse) {
	v.Reset()
	poolTaobaoJstSecretGetAPIResponse.Put(v)
}
