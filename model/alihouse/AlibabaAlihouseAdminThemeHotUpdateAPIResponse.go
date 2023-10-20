package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseadminthemehotupdateAPIResponse 云主题热更新数据集 API返回值
// alibaba.alihouse.admin.theme.hot.update
//
// 云主题更新
type AlibabaalihouseadminthemehotupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseadminthemehotupdateAPIResponseModel
}

// AlibabaalihouseadminthemehotupdateAPIResponseModel is 云主题热更新数据集 成功返回结果
type AlibabaalihouseadminthemehotupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_admin_theme_hot_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihouseadminthemehotupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
