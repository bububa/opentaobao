package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractUserIsloginAPIResponse 校验用户是否已经登录 API返回值
// alibaba.interact.user.islogin
//
// API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
type AlibabaInteractUserIsloginAPIResponse struct {
	model.CommonResponse
	AlibabaInteractUserIsloginAPIResponseModel
}

// AlibabaInteractUserIsloginAPIResponseModel is 校验用户是否已经登录 成功返回结果
type AlibabaInteractUserIsloginAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_user_islogin_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaInteractUserIsloginMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
