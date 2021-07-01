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
type YunosTvpubadminDeviceYksSkillOnlineAPIRequest struct {
    model.Params
    // bot与技能关系表id
    _botSkillRelId   int64
}

// 初始化YunosTvpubadminDeviceYksSkillOnlineAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillOnlineRequest() *YunosTvpubadminDeviceYksSkillOnlineAPIRequest{
    return &YunosTvpubadminDeviceYksSkillOnlineAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.online"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotSkillRelId Setter
// bot与技能关系表id
func (r *YunosTvpubadminDeviceYksSkillOnlineAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
    r._botSkillRelId = _botSkillRelId
    r.Set("bot_skill_rel_id", _botSkillRelId)
    return nil
}

// BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetBotSkillRelId() int64 {
    return r._botSkillRelId
}
