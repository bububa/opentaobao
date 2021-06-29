package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
技能下架 API请求
yunos.tvpubadmin.device.yks.skill.offline

迎客松平台技能下架
*/
type YunosTvpubadminDeviceYksSkillOfflineRequest struct {
    model.Params
    // bot与skill关系表id
    _botSkillRelId   int64
}

// 初始化YunosTvpubadminDeviceYksSkillOfflineRequest对象
func NewYunosTvpubadminDeviceYksSkillOfflineRequest() *YunosTvpubadminDeviceYksSkillOfflineRequest{
    return &YunosTvpubadminDeviceYksSkillOfflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.offline"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotSkillRelId Setter
// bot与skill关系表id
func (r *YunosTvpubadminDeviceYksSkillOfflineRequest) SetBotSkillRelId(_botSkillRelId int64) error {
    r._botSkillRelId = _botSkillRelId
    r.Set("bot_skill_rel_id", _botSkillRelId)
    return nil
}

// BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillOfflineRequest) GetBotSkillRelId() int64 {
    return r._botSkillRelId
}
