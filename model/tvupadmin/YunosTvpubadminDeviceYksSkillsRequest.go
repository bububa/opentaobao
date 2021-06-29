package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备id获取技能列表 API请求
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

// 初始化YunosTvpubadminDeviceYksSkillsRequest对象
func NewYunosTvpubadminDeviceYksSkillsRequest() *YunosTvpubadminDeviceYksSkillsRequest{
    return &YunosTvpubadminDeviceYksSkillsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skills"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotId Setter
// 设备id
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetBotId(botId int64) error {
    r.botId = botId
    r.Set("bot_id", botId)
    return nil
}

// BotId Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetBotId() int64 {
    return r.botId
}
// DeletToken Setter
// 1234
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetDeletToken(deletToken int64) error {
    r.deletToken = deletToken
    r.Set("delet_token", deletToken)
    return nil
}

// DeletToken Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetDeletToken() int64 {
    return r.deletToken
}
// PageIndex Setter
// 当前页
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// PageSize Setter
// 分页单位
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetPageSize() int64 {
    return r.pageSize
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetSkillId() int64 {
    return r.skillId
}
