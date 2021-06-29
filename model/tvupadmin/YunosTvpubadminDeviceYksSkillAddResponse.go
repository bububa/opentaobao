package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加技能 APIResponse
yunos.tvpubadmin.device.yks.skill.add

添加技能
*/
type YunosTvpubadminDeviceYksSkillAddAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillAddResponse
}

type YunosTvpubadminDeviceYksSkillAddResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
