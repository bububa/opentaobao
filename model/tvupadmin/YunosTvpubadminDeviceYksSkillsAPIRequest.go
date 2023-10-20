package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillsAPIRequest 根据设备id获取技能列表 API请求
// yunos.tvpubadmin.device.yks.skills
//
// 根据设备id获取技能列表
type YunosTvpubadminDeviceYksSkillsAPIRequest struct {
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

// NewYunosTvpubadminDeviceYksSkillsRequest 初始化YunosTvpubadminDeviceYksSkillsAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillsRequest() *YunosTvpubadminDeviceYksSkillsAPIRequest {
	return &YunosTvpubadminDeviceYksSkillsAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) Reset() {
	r._skillId = 0
	r._botId = 0
	r._deletToken = 0
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skills"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetBotId is BotId Setter
// 设备id
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetBotId(_botId int64) error {
	r._botId = _botId
	r.Set("bot_id", _botId)
	return nil
}

// GetBotId BotId Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetBotId() int64 {
	return r._botId
}

// SetDeletToken is DeletToken Setter
// 1234
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetDeletToken(_deletToken int64) error {
	r._deletToken = _deletToken
	r.Set("delet_token", _deletToken)
	return nil
}

// GetDeletToken DeletToken Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetDeletToken() int64 {
	return r._deletToken
}

// SetPageIndex is PageIndex Setter
// 当前页
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页单位
func (r *YunosTvpubadminDeviceYksSkillsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminDeviceYksSkillsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolYunosTvpubadminDeviceYksSkillsAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceYksSkillsRequest()
	},
}

// GetYunosTvpubadminDeviceYksSkillsRequest 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillsAPIRequest
func GetYunosTvpubadminDeviceYksSkillsAPIRequest() *YunosTvpubadminDeviceYksSkillsAPIRequest {
	return poolYunosTvpubadminDeviceYksSkillsAPIRequest.Get().(*YunosTvpubadminDeviceYksSkillsAPIRequest)
}

// ReleaseYunosTvpubadminDeviceYksSkillsAPIRequest 将 YunosTvpubadminDeviceYksSkillsAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillsAPIRequest(v *YunosTvpubadminDeviceYksSkillsAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillsAPIRequest.Put(v)
}
