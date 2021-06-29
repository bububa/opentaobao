package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松技能上架接口 API请求
yunos.tvpubadmin.device.yks.skill.online

迎客松技能上架接口
*/
type YunosTvpubadminDeviceYksSkillOnlineRequest struct {
    model.Params
    // bot与技能关系表id
    _botSkillRelId   int64
}

// 初始化YunosTvpubadminDeviceYksSkillOnlineRequest对象
func NewYunosTvpubadminDeviceYksSkillOnlineRequest() *YunosTvpubadminDeviceYksSkillOnlineRequest{
    return &YunosTvpubadminDeviceYksSkillOnlineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillOnlineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.online"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillOnlineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotSkillRelId Setter
// bot与技能关系表id
func (r *YunosTvpubadminDeviceYksSkillOnlineRequest) SetBotSkillRelId(_botSkillRelId int64) error {
    r._botSkillRelId = _botSkillRelId
    r.Set("bot_skill_rel_id", _botSkillRelId)
    return nil
}

// BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillOnlineRequest) GetBotSkillRelId() int64 {
    return r._botSkillRelId
}
