package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretGetAPIResponse 获取TOP通道解密秘钥 API返回值
// taobao.top.secret.get
//
// top sdk通过api获取对应解密秘钥
type TaobaoTopSecretGetAPIResponse struct {
	model.CommonResponse
	TaobaoTopSecretGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopSecretGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopSecretGetAPIResponseModel).Reset()
}

// TaobaoTopSecretGetAPIResponseModel is 获取TOP通道解密秘钥 成功返回结果
type TaobaoTopSecretGetAPIResponseModel struct {
	XMLName xml.Name `xml:"top_secret_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 秘钥值
	Secret string `json:"secret,omitempty" xml:"secret,omitempty"`
	// app配置信息
	AppConfig string `json:"app_config,omitempty" xml:"app_config,omitempty"`
	// 下次更新秘钥间隔，单位（秒）
	Interval int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 最长有效期，容灾使用，单位（秒）
	MaxInterval int64 `json:"max_interval,omitempty" xml:"max_interval,omitempty"`
	// 秘钥版本号
	SecretVersion int64 `json:"secret_version,omitempty" xml:"secret_version,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopSecretGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Secret = ""
	m.AppConfig = ""
	m.Interval = 0
	m.MaxInterval = 0
	m.SecretVersion = 0
}

var poolTaobaoTopSecretGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopSecretGetAPIResponse)
	},
}

// GetTaobaoTopSecretGetAPIResponse 从 sync.Pool 获取 TaobaoTopSecretGetAPIResponse
func GetTaobaoTopSecretGetAPIResponse() *TaobaoTopSecretGetAPIResponse {
	return poolTaobaoTopSecretGetAPIResponse.Get().(*TaobaoTopSecretGetAPIResponse)
}

// ReleaseTaobaoTopSecretGetAPIResponse 将 TaobaoTopSecretGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopSecretGetAPIResponse(v *TaobaoTopSecretGetAPIResponse) {
	v.Reset()
	poolTaobaoTopSecretGetAPIResponse.Put(v)
}
