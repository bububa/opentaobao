package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillOfflineAPIRequest
技能下架 API请求
yunos.tvpubadmin.device.yks.skill.offline

迎客松平台技能下架 */
type YunosTvpubadminDeviceYksSkillOfflineAPIRequest struct {
	model.Params
	// bot与skill关系表id
	_botSkillRelId int64
}

// New
