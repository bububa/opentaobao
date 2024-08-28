package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTaskCreate 聚石塔短信任务创建接口
// taobao.jst.sms.task.create
//
// 聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
func TaobaoJstSmsTaskCreate(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsTaskCreateAPIRequest, resp *jst.TaobaoJstSmsTaskCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
