package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV互动应用活动注册服务 APIResponse
alibaba.interact.activity.register

为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务
*/
type AlibabaInteractActivityRegisterAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_activity_register_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动注册成功，将活动注册后的ID和h5链接返回给调用方
    
    RegisterSucessInfo   *AllsparkResult `json:"register_sucess_info,omitempty" xml:"