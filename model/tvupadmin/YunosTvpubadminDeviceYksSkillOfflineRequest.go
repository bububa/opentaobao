package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
技能下架 APIRequest
yunos.tvpubadmin.device.yks.skill.offline

迎客松平台技能下架
*/
type YunosTvpubadminDeviceYksSkillOfflineRequest struct {
    model.Params

    // bot与skill关系表id
    botSkillRelId   int64 

}

func NewYunosTvpubadminDeviceYksSkillOfflineRequest() *YunosTvpubadminDeviceYksSkillOfflineRequest{
    return &YunosTvpubadminDeviceYksSkillOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksSkillOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.offline"
}

func (r YunosTvpubadminDeviceYksSkillOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceYksSkillOfflineRequest) SetBotSkillRelId(botSkillRelId int64) error {
    r.botSkillRelId = botSkillRelId
    r.Set("bot_skill_rel_id", botSkillRelId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillOfflineRequest) GetBotSkillRelId() int64 {
    return r.botSkillRelId
}

