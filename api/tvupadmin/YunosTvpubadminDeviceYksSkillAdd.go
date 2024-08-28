package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillAdd 添加技能
// yunos.tvpubadmin.device.yks.skill.add
//
// 添加技能
func YunosTvpubadminDeviceYksSkillAdd(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillAddAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
