package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillOnlineAPIRequest
迎客松技能上架接口 API请求
yunos.tvpubadmin.device.yks.skill.online

迎客松技能上架接口 */
type YunosTvpubadminDeviceYksSkillOnlineAPIRequest struct {
	model.Params
	// bot与技能关系表id
	_botSkillRelId int64
}

// New
