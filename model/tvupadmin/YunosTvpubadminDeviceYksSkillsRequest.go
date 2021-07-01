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
type YunosTvpubadminDeviceYksSkillsAPIRequest struct {
    model.Params
    // 设备id
    _botId   int64
    // 1234
    _deletToken   int64
    // 当前页
    _pageIndex   int64
    // 分页单位
    _pageSize   int64
    // 技能id
    _skillId   int64
}

// 初始化YunosTvpubadminDeviceYksSkillsAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillsRequest() *YunosTvpubadminDeviceYksSkillsAPIRequest{
    return &YunosTvpubadminDeviceYksSkillsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skills"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotId Setter
// 设备id
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetBotId(_botId int64) error {
    r._botId = _botId
    r.Set("bot_id", _botId)
    return nil
}

// BotId Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetBotId() int64 {
    return r._botId
}
// DeletToken Setter
// 1234
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetDeletToken(_deletToken int64) error {
    r._deletToken = _deletToken
    r.Set("delet_token", _deletToken)
    return nil
}

// DeletToken Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetDeletToken() int64 {
    return r._deletToken
}
// PageIndex Setter
// 当前页
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 分页单位
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetSkillId() int64 {
    return r._skillId
}
