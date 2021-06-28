package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件会员判断 APIResponse
tmall.device.member.identity.get

用来识别该用户是否是商家会员·
*/
type TmallDeviceMemberIdentityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_device_member_identity_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallDeviceMemberIdentityGetResultDto `json:"result,omitempty" xml:"