package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
技能删除 APIRequest
yunos.tvpubadmin.device.yks.skill.delete

删除技能
*/
type YunosTvpubadminDeviceYksSkillDeleteRequest struct {
    model.Params

    // bot与技能关系表id
    botSkillRelId   int64 

    // 技能id
    skillId   int64 

}

func NewYunosTvpubadminDeviceYksSkillDeleteRequest() *YunosTvpubadminDeviceYksSkillDeleteRequest{
    return &YunosTvpubadminDeviceYksSkillDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.delete"
}

func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceYksSkillDeleteRequest) SetBotSkillRelId(botSkillRelId int64) error {
    r.botSkillRelId = botSkillRelId
    r.Set("bot_skill_rel_id", botSkillRelId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetBotSkillRelId() int64 {
    return r.botSkillRelId
}

func (r *YunosTvpubadminDeviceYksSkillDeleteRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillDeleteRequest) GetSkillId() int64 {
    return r.skillId
}

