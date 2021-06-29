package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id获取技能列表 API返回值 
yunos.tvpubadmin.device.yks.skills

根据设备id获取技能列表
*/
type YunosTvpubadminDeviceYksSkillsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksSkillsResponse
}

// 根据设备id获取技能列表 成功返回结果
type YunosTvpubadminDeviceYksSkillsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skills_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
