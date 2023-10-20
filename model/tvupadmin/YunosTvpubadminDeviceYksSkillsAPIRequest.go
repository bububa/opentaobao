package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskillsAPIRequest 根据设备id获取技能列表 API请求
// yunos.tvpubadmin.device.yks.skills
//
// 根据设备id获取技能列表
type YunostvpubadmindeviceyksskillsAPIRequest struct {
	model.Params
	// 技能id
	_skillId int64
	// 设备id
	_botId int64
	// 1234
	_deletToken int64
	// 当前页
	_pageIndex int64
	// 分页单位
	_pageSize int64
}

// NewYunostvpubadmindeviceyksskillsRequest 初始化YunostvpubadmindeviceyksskillsAPIRequest对象
func NewYunostvpubadmindeviceyksskillsRequest() *YunostvpubadmindeviceyksskillsAPIRequest {
	return &YunostvpubadmindeviceyksskillsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skills"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunostvpubadmindeviceyksskillsAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetBotId is BotId Setter
// 设备id
func (r *YunostvpubadmindeviceyksskillsAPIRequest) SetBotId(_botId int64) error {
	r._botId = _botId
	r.Set("bot_id", _botId)
	return nil
}

// GetBotId BotId Getter
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetBotId() int64 {
	return r._botId
}

// SetDeletToken is DeletToken Setter
// 1234
func (r *YunostvpubadmindeviceyksskillsAPIRequest) SetDeletToken(_deletToken int64) error {
	r._deletToken = _deletToken
	r.Set("delet_token", _deletToken)
	return nil
}

// GetDeletToken DeletToken Getter
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetDeletToken() int64 {
	return r._deletToken
}

// SetPageIndex is PageIndex Setter
// 当前页
func (r *YunostvpubadmindeviceyksskillsAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页单位
func (r *YunostvpubadmindeviceyksskillsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmindeviceyksskillsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
