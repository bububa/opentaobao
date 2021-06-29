package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加技能 APIRequest
yunos.tvpubadmin.device.yks.skill.add

添加技能
*/
type YunosTvpubadminDeviceYksSkillAddRequest struct {
    model.Params

    // 技能id
    skillId   int64 

    // 设备id
    botId   int64 

    // 技能名称
    name   string 

    // 图片地址
    iconImageUrl   string 

}

func NewYunosTvpubadminDeviceYksSkillAddRequest() *YunosTvpubadminDeviceYksSkillAddRequest{
    return &YunosTvpubadminDeviceYksSkillAddRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksSkillAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.add"
}

func (r YunosTvpubadminDeviceYksSkillAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillAddRequest) GetSkillId() int64 {
    return r.skillId
}

func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetBotId(botId int64) error {
    r.botId = botId
    r.Set("bot_id", botId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillAddRequest) GetBotId() int64 {
    return r.botId
}

func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillAddRequest) GetName() string {
    return r.name
}

func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetIconImageUrl(iconImageUrl string) error {
    r.iconImageUrl = iconImageUrl
    r.Set("icon_image_url", iconImageUrl)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillAddRequest) GetIconImageUrl() string {
    return r.iconImageUrl
}

