package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsBotsSkilsList 对外设备获取技能列表
// alibaba.ailabs.bots.skils.list
//
// 获取ai开放平台技能列表
func AlibabaAilabsBotsSkilsList(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsBotsSkilsListAPIRequest, resp *tmallgenie.AlibabaAilabsBotsSkilsListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
