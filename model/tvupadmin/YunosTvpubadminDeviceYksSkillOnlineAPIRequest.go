package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillOnlineAPIRequest 迎客松技能上架接口 API请求
// yunos.tvpubadmin.device.yks.skill.online
//
// 迎客松技能上架接口
type YunosTvpubadminDeviceYksSkillOnlineAPIRequest struct {
	model.Params
	// bot与技能关系表id
	_botSkillRelId int64
}

// NewYunosTvpubadminDeviceYksSkillOnlineRequest 初始化YunosTvpubadminDeviceYksSkillOnlineAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillOnlineRequest() *YunosTvpubadminDeviceYksSkillOnlineAPIRequest {
	return &YunosTvpubadminDeviceYksSkillOnlineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceYksSkillOnlineAPIRequest) Reset() {
	r._botSkillRelId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBotSkillRelId is BotSkillRelId Setter
// bot与技能关系表id
func (r *YunosTvpubadminDeviceYksSkillOnlineAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// GetBotSkillRelId BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillOnlineAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}

var poolYunosTvpubadminDeviceYksSkillOnlineAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceYksSkillOnlineRequest()
	},
}

// GetYunosTvpubadminDeviceYksSkillOnlineRequest 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillOnlineAPIRequest
func GetYunosTvpubadminDeviceYksSkillOnlineAPIRequest() *YunosTvpubadminDeviceYksSkillOnlineAPIRequest {
	return poolYunosTvpubadminDeviceYksSkillOnlineAPIRequest.Get().(*YunosTvpubadminDeviceYksSkillOnlineAPIRequest)
}

// ReleaseYunosTvpubadminDeviceYksSkillOnlineAPIRequest 将 YunosTvpubadminDeviceYksSkillOnlineAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillOnlineAPIRequest(v *YunosTvpubadminDeviceYksSkillOnlineAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillOnlineAPIRequest.Put(v)
}
