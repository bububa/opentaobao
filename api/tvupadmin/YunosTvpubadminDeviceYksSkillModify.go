package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillModify 修改技能
// yunos.tvpubadmin.device.yks.skill.modify
//
// 修改技能
func YunosTvpubadminDeviceYksSkillModify(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
