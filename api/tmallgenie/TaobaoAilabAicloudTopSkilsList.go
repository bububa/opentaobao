package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopSkilsList 获取硬件平台设备下挂载的技能列表
// taobao.ailab.aicloud.top.skils.list
//
// 提供给在硬件平台接入设备的技能列表
func TaobaoAilabAicloudTopSkilsList(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopSkilsListAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopSkilsListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
