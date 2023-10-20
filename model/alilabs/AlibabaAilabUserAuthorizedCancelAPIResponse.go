package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserAuthorizedCancelAPIResponse 取消账号授权 API返回值
// alibaba.ailab.user.authorized.cancel
//
// 三方用户取消授权给天猫精灵用户
type AlibabaAilabUserAuthorizedCancelAPIResponse struct {
	model.CommonResponse
	AlibabaAilabUserAuthorizedCancelAPIResponseModel
}

// AlibabaAilabUserAuthorizedCancelAPIResponseModel is 取消账号授权 成功返回结果
type AlibabaAilabUserAuthorizedCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailab_user_authorized_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误中文描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 返回码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
