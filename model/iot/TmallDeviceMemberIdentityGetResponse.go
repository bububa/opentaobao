package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件会员判断 APIResponse
tmall.device.member.identity.get

用来识别该用户是否是商家会员·
*/
type TmallDeviceMemberIdentityGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallDeviceMemberIdentityGetResponse `json:"tmall_device_member_identity_get_response,omitempty"` 
    TmallDeviceMemberIdentityGetResponse
}

/* model for simplify = false
type TmallDeviceMemberIdentityGetResponse struct {

    // result
    
    Result  *struct {
        TmallDeviceMemberIdentityGetResultDto  *TmallDeviceMemberIdentityGetResultDto `json:"tmall_device_member_identity_get_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallDeviceMemberIdentityGetResponse struct {

    // result
    Result   *TmallDeviceMemberIdentityGetResultDto `json:"result,omitempty"`

}
