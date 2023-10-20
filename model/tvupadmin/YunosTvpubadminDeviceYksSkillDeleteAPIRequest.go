package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskilldeleteAPIRequest 技能删除 API请求
// yunos.tvpubadmin.device.yks.skill.delete
//
// 删除技能
type YunostvpubadmindeviceyksskilldeleteAPIRequest struct {
	model.Params
	// bot与技能关系表id
	_botSkillRelId int64
	// 技能id
	_skillId int64
}

// NewYunostvpubadmindeviceyksskilldeleteRequest 初始化YunostvpubadmindeviceyksskilldeleteAPIRequest对象
func NewYunostvpubadmindeviceyksskilldeleteRequest() *YunostvpubadmindeviceyksskilldeleteAPIRequest {
	return &YunostvpubadmindeviceyksskilldeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceyksskilldeleteAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.skill.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceyksskilldeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceyksskilldeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBotSkillRelId is BotSkillRelId Setter
// bot与技能关系表id
func (r *YunostvpubadmindeviceyksskilldeleteAPIRequest) SetBotSkillRelId(_botSkillRelId int64) error {
	r._botSkillRelId = _botSkillRelId
	r.Set("bot_skill_rel_id", _botSkillRelId)
	return nil
}

// GetBotSkillRelId BotSkillRelId Getter
func (r YunostvpubadmindeviceyksskilldeleteAPIRequest) GetBotSkillRelId() int64 {
	return r._botSkillRelId
}

// SetSkillId is SkillId Setter
// 技能id
func (r *YunostvpubadmindeviceyksskilldeleteAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r YunostvpubadmindeviceyksskilldeleteAPIRequest) GetSkillId() int64 {
	return r._skillId
}
