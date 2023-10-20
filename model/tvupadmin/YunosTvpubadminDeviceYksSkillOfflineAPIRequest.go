package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskillofflineAPIRequest 技能下架 API请求
// yunos.tvpubadmin.device.yks.skill.offline
//
// 迎客松平台技能下架
type YunostvpubadmindeviceyksskillofflineAPIRequest struct {
	model.Params
	// bot与skill关系表id
	_botSkillRelId int64
}

// NewYunostvpubadmindeviceyksskillofflineRequest 初始化YunostvpubadmindeviceyksskillofflineAPIRequest对象
func NewYunostvpubadmindeviceyksskillofflineRequest() *YunostvpubadmindeviceyksskillofflineAPIRequest {
	return &YunostvpubadmindeviceyksskillofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceyksskillofflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceyksskillofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceyksskillofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBotSkillRelId is BotSkillRelId Setter
// bot与skill关系表id
func (r *YunostvpubadmindeviceyksskillofflineAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// GetBotSkillRelId BotSkillRelId Getter
func (r YunostvpubadmindeviceyksskillofflineAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}
