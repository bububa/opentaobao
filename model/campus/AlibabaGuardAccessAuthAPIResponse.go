package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGuardAccessAuthAPIResponse 鉴权 API返回值
// alibaba.guard.access.auth
//
// 刷卡鉴权
type AlibabaGuardAccessAuthAPIResponse struct {
	model.CommonResponse
	AlibabaGuardAccessAuthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaGuardAccessAuthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaGuardAccessAuthAPIResponseModel).Reset()
}

// AlibabaGuardAccessAuthAPIResponseModel is 鉴权 成功返回结果
type AlibabaGuardAccessAuthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_guard_access_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaGuardAccessAuthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaGuardAccessAuthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGuardAccessAuthAPIResponse)
	},
}

// GetAlibabaGuardAccessAuthAPIResponse 从 sync.Pool 获取 AlibabaGuardAccessAuthAPIResponse
func GetAlibabaGuardAccessAuthAPIResponse() *AlibabaGuardAccessAuthAPIResponse {
	return poolAlibabaGuardAccessAuthAPIResponse.Get().(*AlibabaGuardAccessAuthAPIResponse)
}

// ReleaseAlibabaGuardAccessAuthAPIResponse 将 AlibabaGuardAccessAuthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaGuardAccessAuthAPIResponse(v *AlibabaGuardAccessAuthAPIResponse) {
	v.Reset()
	poolAlibabaGuardAccessAuthAPIResponse.Put(v)
}
