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
    Response *TaobaoMiniappEleuserPhoneGetResponse `json:"taobao_miniapp_eleuser_phone_get_response,omitempty"`
}

type TaobaoMiniappEleuserPhoneGetResponse struct {

    // 返回对象
    Result   *EleUicInfo `json:"result,omitempty"`

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

}
