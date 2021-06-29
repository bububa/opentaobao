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
type YunosTvpubadminDeviceYksSkillDeleteRequest struct {
    model.Params
    // bot与技能关系表id
    _botSkillRelId   int64
    // 技能id
    _skillId   int64
}

// 初始化YunosTvpubadminDeviceYksSkillDeleteRequest对象
func NewYunosTvpubadminDeviceYksSkillDeleteRequest() *YunosTvpubadminDeviceYksSkillDeleteRequest{
    return &YunosTvpubadminDeviceYksSkillDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.delete"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotSkillRelId Setter
// bot与技能关系表id
func (r *YunosTvpubadminDeviceYksSkillDeleteRequest) SetBotSkillRelId(_botSkillRelId int64) error {
    r._botSkillRelId = _botSkillRelId
    r.Set("bot_skill_rel_id", _botSkillRelId)
    return nil
}

// BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetBotSkillRelId() int64 {
    return r._botSkillRelId
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillDeleteRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetSkillId() int64 {
    return r._skillId
}
