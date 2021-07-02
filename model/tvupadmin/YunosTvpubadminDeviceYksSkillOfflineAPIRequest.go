package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BotSkillRelId Setter
// bot与skill关系表id
func (r *YunosTvpubadminDeviceYksSkillOfflineAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// Get BotSkillRelId Getter
func (r YunosTvpubadminDeviceYksSkillOfflineAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}
