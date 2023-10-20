package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseAdminThemeUpdateAPIResponse 云主题更新 API返回值
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
type AlibabaAlihouseAdminThemeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseAdminThemeUpdateAPIResponseModel
}

// AlibabaAlihouseAdminThemeUpdateAPIResponseModel is 云主题更新 成功返回结果
type AlibabaAlihouseAdminThemeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseAdminThemeUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
