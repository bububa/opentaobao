package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加技能 API返回值 
yunos.tvpubadmin.device.yks.skill.add

添加技能
*/
type YunosTvpubadminDeviceYksSkillAddAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillAddResponse
}

// 添加技能 成功返回结果
type YunosTvpubadminDeviceYksSkillAddResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}