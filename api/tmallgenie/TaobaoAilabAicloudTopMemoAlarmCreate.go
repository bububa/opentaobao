package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoAlarmCreate 天猫精灵闹钟创建
// taobao.ailab.aicloud.top.memo.alarm.create
//
// 天猫精灵闹钟创建
func TaobaoAilabAicloudTopMemoAlarmCreate(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
