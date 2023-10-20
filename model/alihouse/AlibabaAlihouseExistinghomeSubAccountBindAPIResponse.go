package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeSubAccountBindAPIResponse 子账号入驻 API返回值
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
type AlibabaAlihouseExistinghomeSubAccountBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeSubAccountBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeSubAccountBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeSubAccountBindAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeSubAccountBindAPIResponseModel is 子账号入驻 成功返回结果
type AlibabaAlihouseExistinghomeSubAccountBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_sub_account_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihouseExistinghomeSubAccountBindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeSubAccountBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeSubAccountBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeSubAccountBindAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeSubAccountBindAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeSubAccountBindAPIResponse
func GetAlibabaAlihouseExistinghomeSubAccountBindAPIResponse() *AlibabaAlihouseExistinghomeSubAccountBindAPIResponse {
	return poolAlibabaAlihouseExistinghomeSubAccountBindAPIResponse.Get().(*AlibabaAlihouseExistinghomeSubAccountBindAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeSubAccountBindAPIResponse 将 AlibabaAlihouseExistinghomeSubAccountBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeSubAccountBindAPIResponse(v *AlibabaAlihouseExistinghomeSubAccountBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeSubAccountBindAPIResponse.Put(v)
}
