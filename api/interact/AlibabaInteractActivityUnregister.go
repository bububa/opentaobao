package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractactivityunregister ISV互动应用活动关闭服务
// alibaba.interact.activity.unregister
//
// 卖家在ISV互动应用中设置的活动主动关闭的服务
func Alibabainteractactivityunregister(clt *core.SDKClient, req *interact.AlibabainteractactivityunregisterAPIRequest, session string) (*interact.AlibabainteractactivityunregisterAPIResponse, error) {
	var resp interact.AlibabainteractactivityunregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
