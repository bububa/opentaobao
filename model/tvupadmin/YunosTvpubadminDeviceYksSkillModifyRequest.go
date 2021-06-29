package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改技能 API请求
yunos.tvpubadmin.device.yks.skill.modify

修改技能
*/
type YunosTvpubadminDeviceYksSkillModifyRequest struct {
    model.Params
    // 技能id
    skillId   int64
    // 图片地址
    iconImageUrl   string
    // 技能名称
    name   string
}

// 初始化YunosTvpubadminDeviceYksSkillModifyRequest对象
func NewYunosTvpubadminDeviceYksSkillModifyRequest() *YunosTvpubadminDeviceYksSkillModifyRequest{
    return &YunosTvpubadminDeviceYksSkillModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.modify"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillModifyRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetSkillId() int64 {
    return r.skillId
}
// IconImageUrl Setter
// 图片地址
func (r *YunosTvpubadminDeviceYksSkillModifyRequest) SetIconImageUrl(iconImageUrl string) error {
    r.iconImageUrl = iconImageUrl
    r.Set("icon_image_url", iconImageUrl)
    return nil
}

// IconImageUrl Getter
func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetIconImageUrl() string {
    return r.iconImageUrl
}
// Name Setter
// 技能名称
func (r *YunosTvpubadminDeviceYksSkillModifyRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetName() string {
    return r.name
}
