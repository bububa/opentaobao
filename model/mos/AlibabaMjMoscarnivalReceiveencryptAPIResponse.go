package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMoscarnivalReceiveencryptAPIResponse 根据加密手机号领券 API返回值
// alibaba.mj.moscarnival.receiveencrypt
//
// 根据加密手机号领券
type AlibabaMjMoscarnivalReceiveencryptAPIResponse struct {
	model.CommonResponse
	AlibabaMjMoscarnivalReceiveencryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMoscarnivalReceiveencryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMoscarnivalReceiveencryptAPIResponseModel).Reset()
}

// AlibabaMjMoscarnivalReceiveencryptAPIResponseModel is 根据加密手机号领券 成功返回结果
type AlibabaMjMoscarnivalReceiveencryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_moscarnival_receiveencrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMjMoscarnivalReceiveencryptResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMoscarnivalReceiveencryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMjMoscarnivalReceiveencryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMoscarnivalReceiveencryptAPIResponse)
	},
}

// GetAlibabaMjMoscarnivalReceiveencryptAPIResponse 从 sync.Pool 获取 AlibabaMjMoscarnivalReceiveencryptAPIResponse
func GetAlibabaMjMoscarnivalReceiveencryptAPIResponse() *AlibabaMjMoscarnivalReceiveencryptAPIResponse {
	return poolAlibabaMjMoscarnivalReceiveencryptAPIResponse.Get().(*AlibabaMjMoscarnivalReceiveencryptAPIResponse)
}

// ReleaseAlibabaMjMoscarnivalReceiveencryptAPIResponse 将 AlibabaMjMoscarnivalReceiveencryptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMoscarnivalReceiveencryptAPIResponse(v *AlibabaMjMoscarnivalReceiveencryptAPIResponse) {
	v.Reset()
	poolAlibabaMjMoscarnivalReceiveencryptAPIResponse.Put(v)
}
