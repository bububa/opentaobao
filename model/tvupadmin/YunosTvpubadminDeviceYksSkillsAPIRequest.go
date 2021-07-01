package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillsAPIRequest
根据设备id获取技能列表 API请求
yunos.tvpubadmin.device.yks.skills

根据设备id获取技能列表 */
type YunosTvpubadminDeviceYksSkillsAPIRequest struct {
	model.Params
	// 设备id
	_botId int64
	// 1234
	_deletToken int64
	// 当前页
	_pageIndex int64
	// 分页单位
	_pageSize int64
	// 技能id
	_skillId int64
}

// New
