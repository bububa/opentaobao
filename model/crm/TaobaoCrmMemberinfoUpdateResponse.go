package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑会员资料 APIResponse
taobao.crm.memberinfo.update

编辑会员的基本资料，接口返回会员信息修改是否成功
*/
type TaobaoCrmMemberinfoUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_memberinfo_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 会员信息修改是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"