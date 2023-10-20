package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillOnline 迎客松技能上架接口
// yunos.tvpubadmin.device.yks.skill.online
//
// 迎客松技能上架接口
func YunosTvpubadminDeviceYksSkillOnline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillOnlineAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillOnlineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
