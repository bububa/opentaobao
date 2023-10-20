package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeCreateAPIResponse 创建云主题 API返回值
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
type AlibabaAlihouseAdminThemeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseAdminThemeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseAdminThemeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseAdminThemeCreateAPIResponseModel).Reset()
}

// AlibabaAlihouseAdminThemeCreateAPIResponseModel is 创建云主题 成功返回结果
type AlibabaAlihouseAdminThemeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseAdminThemeCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseAdminThemeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseAdminThemeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeCreateAPIResponse)
	},
}

// GetAlibabaAlihouseAdminThemeCreateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseAdminThemeCreateAPIResponse
func GetAlibabaAlihouseAdminThemeCreateAPIResponse() *AlibabaAlihouseAdminThemeCreateAPIResponse {
	return poolAlibabaAlihouseAdminThemeCreateAPIResponse.Get().(*AlibabaAlihouseAdminThemeCreateAPIResponse)
}

// ReleaseAlibabaAlihouseAdminThemeCreateAPIResponse 将 AlibabaAlihouseAdminThemeCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeCreateAPIResponse(v *AlibabaAlihouseAdminThemeCreateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeCreateAPIResponse.Put(v)
}
