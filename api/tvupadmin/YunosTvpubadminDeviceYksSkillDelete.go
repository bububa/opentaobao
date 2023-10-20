package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillDelete 技能删除
// yunos.tvpubadmin.device.yks.skill.delete
//
// 删除技能
func YunosTvpubadminDeviceYksSkillDelete(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillDeleteAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
