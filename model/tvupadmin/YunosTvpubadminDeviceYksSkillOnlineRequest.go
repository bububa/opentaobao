package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松技能上架接口 APIRequest
yunos.tvpubadmin.device.yks.skill.online

迎客松技能上架接口
*/
type YunosTvpubadminDeviceYksSkillOnlineRequest struct {
    model.Params

    // bot与技能关系表id
    botSkillRelId   int64 

}

func NewYunosTvpubadminDeviceYksSkillOnlineRequest() *YunosTvpubadminDeviceYksSkillOnlineRequest{
    return &YunosTvpubadminDeviceYksSkillOnlineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksSkillOnlineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.online"
}

func (r YunosTvpubadminDeviceYksSkillOnlineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceYksSkillOnlineRequest) SetBotSkillRelId(botSkillRelId int64) error {
    r.botSkillRelId = botSkillRelId
    r.Set("bot_skill_rel_id", botSkillRelId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillOnlineRequest) GetBotSkillRelId() int64 {
    return r.botSkillRelId
}

