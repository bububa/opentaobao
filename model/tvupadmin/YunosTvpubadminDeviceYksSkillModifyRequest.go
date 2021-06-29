package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改技能 APIRequest
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

func NewYunosTvpubadminDeviceYksSkillModifyRequest() *YunosTvpubadminDeviceYksSkillModifyRequest{
    return &YunosTvpubadminDeviceYksSkillModifyRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.modify"
}

func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceYksSkillModifyRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetSkillId() int64 {
    return r.skillId
}

func (r *YunosTvpubadminDeviceYksSkillModifyRequest) SetIconImageUrl(iconImageUrl string) error {
    r.iconImageUrl = iconImageUrl
    r.Set("icon_image_url", iconImageUrl)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetIconImageUrl() string {
    return r.iconImageUrl
}

func (r *YunosTvpubadminDeviceYksSkillModifyRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillModifyRequest) GetName() string {
    return r.name
}

