package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleUserPermit 用户appkey授权
// alibaba.idle.user.permit
//
// 闲鱼卖家与服务商关系绑定，用于业务数据分发/授权校验/消息通知鉴权
func AlibabaIdleUserPermit(clt *core.SDKClient, req *idle.AlibabaIdleUserPermitAPIRequest, resp *idle.AlibabaIdleUserPermitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
