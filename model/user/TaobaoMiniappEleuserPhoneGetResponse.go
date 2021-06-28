package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 APIResponse
taobao.miniapp.eleuser.phone.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserPhoneGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMiniappEleuserPhoneGetResponse `json:"miniapp_eleuser_phone_get_response,omitempty"` 
    TaobaoMiniappEleuserPhoneGetResponse
}

/* model for simplify = false
type TaobaoMiniappEleuserPhoneGetResponse struct {

    // 返回对象
    
    Result  *struct {
        EleUicInfo  *EleUicInfo `json:"ele_uic_info,omitempty"`
    } `json:"result,omitempty"`
    

    // traceId
    
    TraceId   string `json:"trace_id,omitempty"`
    

}
*/

type TaobaoMiniappEleuserPhoneGetResponse struct {

    // 返回对象
    Result   *EleUicInfo `json:"result,omitempty"`

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

}
