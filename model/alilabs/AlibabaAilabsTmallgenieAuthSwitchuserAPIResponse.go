package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse 切换用户 API返回值
// alibaba.ailabs.tmallgenie.auth.switchuser
//
// 设备切换授权用户
type AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthSwitchuserAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthSwitchuserAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthSwitchuserAPIResponseModel is 切换用户 成功返回结果
type AlibabaAilabsTmallgenieAuthSwitchuserAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_switchuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAilabsTmallgenieAuthSwitchuserResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthSwitchuserAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse
func GetAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse() *AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse 将 AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse(v *AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthSwitchuserAPIResponse.Put(v)
}
