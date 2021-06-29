package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id获取技能列表 APIRequest
yunos.tvpubadmin.device.yks.skills

根据设备id获取技能列表
*/
type YunosTvpubadminDeviceYksSkillsRequest struct {
    model.Params

    // 设备id
    botId   int64 

    // 1234
    deletToken   int64 

    // 当前页
    pageIndex   int64 

    // 分页单位
    pageSize   int64 

    // 技能id
    skillId   int64 

}

func NewYunosTvpubadminDeviceYksSkillsRequest() *YunosTvpubadminDeviceYksSkillsRequest{
    return &YunosTvpubadminDeviceYksSkillsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skills"
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceYksSkillsRequest) SetBotId(botId int64) error {
    r.botId = botId
    r.Set("bot_id", botId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetBotId() int64 {
    return r.botId
}

func (r *YunosTvpubadminDeviceYksSkillsRequest) SetDeletToken(deletToken int64) error {
    r.deletToken = deletToken
    r.Set("delet_token", deletToken)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetDeletToken() int64 {
    return r.deletToken
}

func (r *YunosTvpubadminDeviceYksSkillsRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *YunosTvpubadminDeviceYksSkillsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *YunosTvpubadminDeviceYksSkillsRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r YunosTvpubadminDeviceYksSkillsRequest) GetSkillId() int64 {
    return r.skillId
}

