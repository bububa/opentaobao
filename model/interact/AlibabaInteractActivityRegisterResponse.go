package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ISV互动应用活动注册服务 APIResponse
alibaba.interact.activity.register

为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务
*/
type AlibabaInteractActivityRegisterAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractActivityRegisterResponse `json:"alibaba_interact_activity_register_response,omitempty"` 
    AlibabaInteractActivityRegisterResponse
}

/* model for simplify = false
type AlibabaInteractActivityRegisterResponse struct {

    // 活动注册成功，将活动注册后的ID和h5链接返回给调用方
    
    RegisterSucessInfo  *struct {
        AllsparkResult  *AllsparkResult `json:"allspark_result,omitempty"`
    } `json:"register_sucess_info,omitempty"`
    

}
*/

type AlibabaInteractActivityRegisterResponse struct {

    // 活动注册成功，将活动注册后的ID和h5链接返回给调用方
    RegisterSucessInfo   *AllsparkResult `json:"register_sucess_info,omitempty"`

}
