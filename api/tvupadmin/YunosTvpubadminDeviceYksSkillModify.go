package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillModify 修改技能
// yunos.tvpubadmin.device.yks.skill.modify
//
// 修改技能
func YunosTvpubadminDeviceYksSkillModify(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
