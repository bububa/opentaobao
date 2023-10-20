package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillOffline 技能下架
// yunos.tvpubadmin.device.yks.skill.offline
//
// 迎客松平台技能下架
func YunosTvpubadminDeviceYksSkillOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
