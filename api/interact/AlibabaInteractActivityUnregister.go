package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractActivityUnregister ISV互动应用活动关闭服务
// alibaba.interact.activity.unregister
//
// 卖家在ISV互动应用中设置的活动主动关闭的服务
func AlibabaInteractActivityUnregister(clt *core.SDKClient, req *interact.AlibabaInteractActivityUnregisterAPIRequest, resp *interact.AlibabaInteractActivityUnregisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
