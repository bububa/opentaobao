package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeUpdateStatusAPIResponse 云主题上下架+删除 API返回值
// alibaba.alihouse.admin.theme.update.status
//
// 云主题上下架+删除
type AlibabaAlihouseAdminThemeUpdateStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseAdminThemeUpdateStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel).Reset()
}

// AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel is 云主题上下架+删除 成功返回结果
type AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_update_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseAdminThemeUpdateStatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseAdminThemeUpdateStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeUpdateStatusAPIResponse)
	},
}

// GetAlibabaAlihouseAdminThemeUpdateStatusAPIResponse 从 sync.Pool 获取 AlibabaAlihouseAdminThemeUpdateStatusAPIResponse
func GetAlibabaAlihouseAdminThemeUpdateStatusAPIResponse() *AlibabaAlihouseAdminThemeUpdateStatusAPIResponse {
	return poolAlibabaAlihouseAdminThemeUpdateStatusAPIResponse.Get().(*AlibabaAlihouseAdminThemeUpdateStatusAPIResponse)
}

// ReleaseAlibabaAlihouseAdminThemeUpdateStatusAPIResponse 将 AlibabaAlihouseAdminThemeUpdateStatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeUpdateStatusAPIResponse(v *AlibabaAlihouseAdminThemeUpdateStatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeUpdateStatusAPIResponse.Put(v)
}
