package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillModifyAPIRequest
修改技能 API请求
yunos.tvpubadmin.device.yks.skill.modify

修改技能 */
type YunosTvpubadminDeviceYksSkillModifyAPIRequest struct {
	model.Params
	// 技能id
	_skillId int64
	// 图片地址
	_iconImageUrl string
	// 技能名称
	_name string
}

// New
