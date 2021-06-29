package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
技能删除 APIResponse
yunos.tvpubadmin.device.yks.skill.delete

删除技能
*/
type YunosTvpubadminDeviceYksSkillDeleteAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillDeleteResponse
}

type YunosTvpubadminDeviceYksSkillDeleteResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
