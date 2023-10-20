package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillOfflineAPIRequest 技能下架 API请求
// yunos.tvpubadmin.device.yks.skill.offline
//
// 迎客松平台技能下架
type YunosTvpubadminDeviceYksSkillOfflineAPIRequest struct {
	model.Params
	// bot与skill关系表id
	_botSkillRelId int64
}

// NewYunosTvpubadminDeviceYksSkillOfflineRequest 初始化YunosTvpubadminDeviceYksSkillOfflineAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillOfflineRequest() *YunosTvpubadminDeviceYksSkillOfflineAPIRequest {
	return &YunosTvpubadminDeviceYksSkillOfflineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceYksSkillOfflineAPIRequest) Reset() {
	r._botSkillRelId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBotSkillRelId is BotSkillRelId Setter
// bot与skill关系表id
func (r *YunosTvpubadminDeviceYksSkillOfflineAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// GetBotSkillRelId BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}

var poolYunosTvpubadminDeviceYksSkillOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceYksSkillOfflineRequest()
	},
}

// GetYunosTvpubadminDeviceYksSkillOfflineRequest 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillOfflineAPIRequest
func GetYunosTvpubadminDeviceYksSkillOfflineAPIRequest() *YunosTvpubadminDeviceYksSkillOfflineAPIRequest {
	return poolYunosTvpubadminDeviceYksSkillOfflineAPIRequest.Get().(*YunosTvpubadminDeviceYksSkillOfflineAPIRequest)
}

// ReleaseYunosTvpubadminDeviceYksSkillOfflineAPIRequest 将 YunosTvpubadminDeviceYksSkillOfflineAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillOfflineAPIRequest(v *YunosTvpubadminDeviceYksSkillOfflineAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillOfflineAPIRequest.Put(v)
}
