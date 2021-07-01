package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
技能删除 API请求
yunos.tvpubadmin.device.yks.skill.delete

删除技能
*/
type YunosTvpubadminDeviceYksSkillDeleteAPIRequest struct {
    model.Params
    // bot与技能关系表id
    _botSkillRelId   int64
    // 技能id
    _skillId   int64
}

// 初始化YunosTvpubadminDeviceYksSkillDeleteAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillDeleteRequest() *YunosTvpubadminDeviceYksSkillDeleteAPIRequest{
    return &YunosTvpubadminDeviceYksSkillDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.delete"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotSkillRelId Setter
// bot与技能关系表id
func (r *YunosTvpubadminDeviceYksSkillDeleteAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
    r._botSkillRelId = _botSkillRelId
    r.Set("bot_skill_rel_id", _botSkillRelId)
    return nil
}

// BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetBotSkillRelId() int64 {
    return r._botSkillRelId
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillDeleteAPIRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetSkillId() int64 {
    return r._skillId
}
