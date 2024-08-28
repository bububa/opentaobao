package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopMessageSend 发送留言
// taobao.ailab.aicloud.top.message.send
//
// 供准入的外部用户实现发送留言功能，APP端发送，设备端读取
func TaobaoAilabAicloudTopMessageSend(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageSendAPIRequest, resp *iot.TaobaoAilabAicloudTopMessageSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
