package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松技能上架接口 APIResponse
yunos.tvpubadmin.device.yks.skill.online

迎客松技能上架接口
*/
type YunosTvpubadminDeviceYksSkillOnlineAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillOnlineResponse
}

type YunosTvpubadminDeviceYksSkillOnlineResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_online_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
