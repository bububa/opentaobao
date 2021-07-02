package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityRegisterAPIResponse ISV互动应用活动注册服务 API返回值
// alibaba.interact.activity.register
//
// 为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务
type AlibabaInteractActivityRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaInteractActivityRegisterAPIResponseModel
}

// AlibabaInteractActivityRegisterAPIResponseModel is ISV互动应用活动注册服务 成功返回结果
type AlibabaInteractActivityRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动注册成功，将活动注册后的ID和h5链接返回给调用方
	RegisterSucessInfo *AllsparkResult `json:"register_sucess_info,omitempty" xml:"register_sucess_info,omitempty"`
}
