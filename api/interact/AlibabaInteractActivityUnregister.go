package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractActivityUnregister
ISV互动应用活动关闭服务
alibaba.interact.activity.unregister

卖家在ISV互动应用中设置的活动主动关闭的服务 */
func AlibabaInteractActivityUnregister(clt *core.SDKClient, req *interact.AlibabaInteractActivityUnregisterAPIRequest, session string) (*interact.AlibabaInteractActivityUnregisterAPIResponse, error) {
	var resp interact.AlibabaInteractActivityUnregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
