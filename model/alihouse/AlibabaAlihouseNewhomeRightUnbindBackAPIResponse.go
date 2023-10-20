package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRightUnbindBackAPIResponse 权限回流-解绑 API返回值
// alibaba.alihouse.newhome.right.unbind.back
//
// 权限回流-解绑
type AlibabaAlihouseNewhomeRightUnbindBackAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeRightUnbindBackAPIResponseModel
}

// AlibabaAlihouseNewhomeRightUnbindBackAPIResponseModel is 权限回流-解绑 成功返回结果
type AlibabaAlihouseNewhomeRightUnbindBackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_right_unbind_back_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlibabaAlihouseNewhomeRightUnbindBackResult `json:"result,omitempty" xml:"result,omitempty"`
}
