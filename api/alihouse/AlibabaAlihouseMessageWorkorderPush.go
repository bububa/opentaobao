package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMessageWorkorderPush 工单消息推送
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
func AlibabaAlihouseMessageWorkorderPush(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseMessageWorkorderPushAPIRequest, resp *alihouse.AlibabaAlihouseMessageWorkorderPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
