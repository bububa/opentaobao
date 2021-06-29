package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改技能 API返回值 
yunos.tvpubadmin.device.yks.skill.modify

修改技能
*/
type YunosTvpubadminDeviceYksSkillModifyAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillModifyResponse
}

// 修改技能 成功返回结果
type YunosTvpubadminDeviceYksSkillModifyResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
