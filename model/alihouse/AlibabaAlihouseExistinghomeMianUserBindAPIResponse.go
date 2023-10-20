package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeMianUserBindAPIResponse 主账号入驻 API返回值
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
type AlibabaAlihouseExistinghomeMianUserBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeMianUserBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeMianUserBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeMianUserBindAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeMianUserBindAPIResponseModel is 主账号入驻 成功返回结果
type AlibabaAlihouseExistinghomeMianUserBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_mian_user_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihouseExistinghomeMianUserBindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeMianUserBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeMianUserBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeMianUserBindAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeMianUserBindAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeMianUserBindAPIResponse
func GetAlibabaAlihouseExistinghomeMianUserBindAPIResponse() *AlibabaAlihouseExistinghomeMianUserBindAPIResponse {
	return poolAlibabaAlihouseExistinghomeMianUserBindAPIResponse.Get().(*AlibabaAlihouseExistinghomeMianUserBindAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeMianUserBindAPIResponse 将 AlibabaAlihouseExistinghomeMianUserBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeMianUserBindAPIResponse(v *AlibabaAlihouseExistinghomeMianUserBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeMianUserBindAPIResponse.Put(v)
}
