package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
技能下架 APIResponse
yunos.tvpubadmin.device.yks.skill.offline

迎客松平台技能下架
*/
type YunosTvpubadminDeviceYksSkillOfflineAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillOfflineResponse
}

type YunosTvpubadminDeviceYksSkillOfflineResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_offline_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
