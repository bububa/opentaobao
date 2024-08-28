package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsTaowaiMsgSend 淘外包裹物流信息走淘宝发包裹状态通知接口
// taobao.logistics.taowai.msg.send
//
// 淘外包裹物流信息走淘宝发包裹状态通知接口
func TaobaoLogisticsTaowaiMsgSend(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoLogisticsTaowaiMsgSendAPIRequest, resp *mtopopen.TaobaoLogisticsTaowaiMsgSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
