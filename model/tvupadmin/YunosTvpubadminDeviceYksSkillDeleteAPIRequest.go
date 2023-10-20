package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillDeleteAPIRequest 技能删除 API请求
// yunos.tvpubadmin.device.yks.skill.delete
//
// 删除技能
type YunosTvpubadminDeviceYksSkillDeleteAPIRequest struct {
	model.Params
	// bot与技能关系表id
	_botSkillRelId int64
	// 技能id
	_skillId int64
}

// NewYunosTvpubadminDeviceYksSkillDeleteRequest 初始化YunosTvpubadminDeviceYksSkillDeleteAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillDeleteRequest() *YunosTvpubadminDeviceYksSkillDeleteAPIRequest {
	return &YunosTvpubadminDeviceYksSkillDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceYksSkillDeleteAPIRequest) Reset() {
	r._botSkillRelId = 0
	r._skillId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBotSkillRelId is BotSkillRelId Setter
// bot与技能关系表id
func (r *YunosTvpubadminDeviceYksSkillDeleteAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// GetBotSkillRelId BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillDeleteAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunosTvpubadminDeviceYksSkillDeleteAPIRequest) GetSkillId() int64 {
	return r._skillId
}

var poolYunosTvpubadminDeviceYksSkillDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceYksSkillDeleteRequest()
	},
}

// GetYunosTvpubadminDeviceYksSkillDeleteRequest 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillDeleteAPIRequest
func GetYunosTvpubadminDeviceYksSkillDeleteAPIRequest() *YunosTvpubadminDeviceYksSkillDeleteAPIRequest {
	return poolYunosTvpubadminDeviceYksSkillDeleteAPIRequest.Get().(*YunosTvpubadminDeviceYksSkillDeleteAPIRequest)
}

// ReleaseYunosTvpubadminDeviceYksSkillDeleteAPIRequest 将 YunosTvpubadminDeviceYksSkillDeleteAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillDeleteAPIRequest(v *YunosTvpubadminDeviceYksSkillDeleteAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillDeleteAPIRequest.Put(v)
}
