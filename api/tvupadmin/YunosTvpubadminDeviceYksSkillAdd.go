package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillAdd 添加技能
// yunos.tvpubadmin.device.yks.skill.add
//
// 添加技能
func YunosTvpubadminDeviceYksSkillAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillAddAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceYksSkillAddAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceYksSkillAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
