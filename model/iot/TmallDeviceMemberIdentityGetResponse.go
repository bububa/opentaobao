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
    Response *TmallDeviceMemberIdentityGetResponse `json:"tmall_device_member_identity_get_response,omitempty"`
}

type TmallDeviceMemberIdentityGetResponse struct {

    // result
    Result   *TmallDeviceMemberIdentityGetResultDto `json:"result,omitempty"`

}
