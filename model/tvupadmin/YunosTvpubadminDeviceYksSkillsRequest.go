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
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetBotId(_botId int64) error {
    r._botId = _botId
    r.Set("bot_id", _botId)
    return nil
}

// BotId Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetBotId() int64 {
    return r._botId
}
// DeletToken Setter
// 1234
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetDeletToken(_deletToken int64) error {
    r._deletToken = _deletToken
    r.Set("delet_token", _deletToken)
    return nil
}

// DeletToken Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetDeletToken() int64 {
    return r._deletToken
}
// PageIndex Setter
// 当前页
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 分页单位
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetPageSize() int64 {
    return r._pageSize
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillsRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillsRequest) GetSkillId() int64 {
    return r._skillId
}
