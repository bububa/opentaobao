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
type YunosTvpubadminDeviceYksSkillModifyAPIRequest struct {
    model.Params
    // 技能id
    _skillId   int64
    // 图片地址
    _iconImageUrl   string
    // 技能名称
    _name   string
}

// 初始化YunosTvpubadminDeviceYksSkillModifyAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillModifyRequest() *YunosTvpubadminDeviceYksSkillModifyAPIRequest{
    return &YunosTvpubadminDeviceYksSkillModifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.modify"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillModifyAPIRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetSkillId() int64 {
    return r._skillId
}
// IconImageUrl Setter
// 图片地址
func (r *YunosTvpubadminDeviceYksSkillModifyAPIRequest) SetIconImageUrl(_iconImageUrl string) error {
    r._iconImageUrl = _iconImageUrl
    r.Set("icon_image_url", _iconImageUrl)
    return nil
}

// IconImageUrl Getter
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetIconImageUrl() string {
    return r._iconImageUrl
}
// Name Setter
// 技能名称
func (r *YunosTvpubadminDeviceYksSkillModifyAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r YunosTvpubadminDeviceYksSkillModifyAPIRequest) GetName() string {
    return r._name
}
