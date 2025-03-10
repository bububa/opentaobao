package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsFengchaoMsgSend 丰巢走淘宝发包裹状态通知接口
// taobao.logistics.fengchao.msg.send
//
// 丰巢走淘宝发包裹状态通知接口
func TaobaoLogisticsFengchaoMsgSend(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoLogisticsFengchaoMsgSendAPIRequest, resp *mtopopen.TaobaoLogisticsFengchaoMsgSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
