package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeHotUpdateAPIResponse 云主题热更新数据集 API返回值
// alibaba.alihouse.admin.theme.hot.update
//
// 云主题更新
type AlibabaAlihouseAdminThemeHotUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseAdminThemeHotUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseAdminThemeHotUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseAdminThemeHotUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseAdminThemeHotUpdateAPIResponseModel is 云主题热更新数据集 成功返回结果
type AlibabaAlihouseAdminThemeHotUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_hot_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseAdminThemeHotUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseAdminThemeHotUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseAdminThemeHotUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseAdminThemeHotUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseAdminThemeHotUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseAdminThemeHotUpdateAPIResponse
func GetAlibabaAlihouseAdminThemeHotUpdateAPIResponse() *AlibabaAlihouseAdminThemeHotUpdateAPIResponse {
	return poolAlibabaAlihouseAdminThemeHotUpdateAPIResponse.Get().(*AlibabaAlihouseAdminThemeHotUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseAdminThemeHotUpdateAPIResponse 将 AlibabaAlihouseAdminThemeHotUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseAdminThemeHotUpdateAPIResponse(v *AlibabaAlihouseAdminThemeHotUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseAdminThemeHotUpdateAPIResponse.Put(v)
}
