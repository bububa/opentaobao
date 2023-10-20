package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthuicuserinfohealthidgetAPIResponse 获取健康id API返回值
// alibaba.alihealth.uic.userinfo.healthid.get
//
// 根据支付宝用户ID获取用户健康ID
type AlibabaalihealthuicuserinfohealthidgetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthuicuserinfohealthidgetAPIResponseModel
}

// AlibabaalihealthuicuserinfohealthidgetAPIResponseModel is 获取健康id 成功返回结果
type AlibabaalihealthuicuserinfohealthidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_uic_userinfo_healthid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
