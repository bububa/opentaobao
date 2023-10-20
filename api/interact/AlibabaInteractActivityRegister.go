package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractActivityRegister ISV互动应用活动注册服务
// alibaba.interact.activity.register
//
// 为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务
func AlibabaInteractActivityRegister(clt *core.SDKClient, req *interact.AlibabaInteractActivityRegisterAPIRequest, resp *interact.AlibabaInteractActivityRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
