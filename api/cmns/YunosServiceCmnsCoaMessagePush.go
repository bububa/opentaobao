package cmns

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaMessagePush 消息推送接口
// yunos.service.cmns.coa.message.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
func YunosServiceCmnsCoaMessagePush(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessagePushAPIRequest, resp *cmns.YunosServiceCmnsCoaMessagePushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
