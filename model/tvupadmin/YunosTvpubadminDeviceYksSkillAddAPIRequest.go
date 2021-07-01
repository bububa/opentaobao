package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillAddAPIRequest
添加技能 API请求
yunos.tvpubadmin.device.yks.skill.add

添加技能 */
type YunosTvpubadminDeviceYksSkillAddAPIRequest struct {
	model.Params
	// 技能id
	_skillId int64
	// 设备id
	_botId int64
	// 技能名称
	_name string
	// 图片地址
	_iconImageUrl string
}

// New
