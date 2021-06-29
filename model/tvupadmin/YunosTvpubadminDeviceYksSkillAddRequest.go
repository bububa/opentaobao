package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加技能 API请求
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

// 初始化YunosTvpubadminDeviceYksSkillAddRequest对象
func NewYunosTvpubadminDeviceYksSkillAddRequest() *YunosTvpubadminDeviceYksSkillAddRequest{
    return &YunosTvpubadminDeviceYksSkillAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillAddRequest) GetSkillId() int64 {
    return r.skillId
}
// BotId Setter
// 设备id
func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetBotId(botId int64) error {
    r.botId = botId
    r.Set("bot_id", botId)
    return nil
}

// BotId Getter
func (r YunosTvpubadminDeviceYksSkillAddRequest) GetBotId() int64 {
    return r.botId
}
// Name Setter
// 技能名称
func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r YunosTvpubadminDeviceYksSkillAddRequest) GetName() string {
    return r.name
}
// IconImageUrl Setter
// 图片地址
func (r *YunosTvpubadminDeviceYksSkillAddRequest) SetIconImageUrl(iconImageUrl string) error {
    r.iconImageUrl = iconImageUrl
    r.Set("icon_image_url", iconImageUrl)
    return nil
}

// IconImageUrl Getter
func (r YunosTvpubadminDeviceYksSkillAddRequest) GetIconImageUrl() string {
    return r.iconImageUrl
}
