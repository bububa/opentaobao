package alihouse

import (
	"encoding/xml"

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

// AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel is 云主题上下架+删除 成功返回结果
type AlibabaAlihouseAdminThemeUpdateStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_update_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseAdminThemeUpdateStatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
