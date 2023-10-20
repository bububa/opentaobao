package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseadminthemecreateAPIResponse 创建云主题 API返回值
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
type AlibabaalihouseadminthemecreateAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseadminthemecreateAPIResponseModel
}

// AlibabaalihouseadminthemecreateAPIResponseModel is 创建云主题 成功返回结果
type AlibabaalihouseadminthemecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihouseadminthemecreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
