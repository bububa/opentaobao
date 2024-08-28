package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoQimenEventsProduce 批量发送奇门事件
// taobao.qimen.events.produce
//
// 批量发送消息
func TaobaoQimenEventsProduce(ctx context.Context, clt *core.SDKClient, req *util.TaobaoQimenEventsProduceAPIRequest, resp *util.TaobaoQimenEventsProduceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
