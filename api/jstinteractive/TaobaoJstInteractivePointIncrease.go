package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractivePointIncrease 互动积分发放接口
// taobao.jst.interactive.point.increase
//
// 向用户发放互动积分
func TaobaoJstInteractivePointIncrease(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointIncreaseAPIRequest, resp *jstinteractive.TaobaoJstInteractivePointIncreaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
