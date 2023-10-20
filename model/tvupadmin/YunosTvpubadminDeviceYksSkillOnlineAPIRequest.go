package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskillonlineAPIRequest 迎客松技能上架接口 API请求
// yunos.tvpubadmin.device.yks.skill.online
//
// 迎客松技能上架接口
type YunostvpubadmindeviceyksskillonlineAPIRequest struct {
	model.Params
	// bot与技能关系表id
	_botSkillRelId int64
}

// NewYunostvpubadmindeviceyksskillonlineRequest 初始化YunostvpubadmindeviceyksskillonlineAPIRequest对象
func NewYunostvpubadmindeviceyksskillonlineRequest() *YunostvpubadmindeviceyksskillonlineAPIRequest {
	return &YunostvpubadmindeviceyksskillonlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceyksskillonlineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceyksskillonlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceyksskillonlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBotSkillRelId is BotSkillRelId Setter
// bot与技能关系表id
func (r *YunostvpubadmindeviceyksskillonlineAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// GetBotSkillRelId BotSkillRelId Getter
func (r YunostvpubadmindeviceyksskillonlineAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}
