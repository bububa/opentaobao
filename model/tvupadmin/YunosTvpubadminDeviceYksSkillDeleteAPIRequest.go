package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillDeleteAPIRequest
技能删除 API请求
yunos.tvpubadmin.device.yks.skill.delete

删除技能 */
type YunosTvpubadminDeviceYksSkillDeleteAPIRequest struct {
	model.Params
	// bot与技能关系表id
	_botSkillRelId int64
	// 技能id
	_skillId int64
}

// New
