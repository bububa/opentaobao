package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractivePointDecrease 互动积分扣减接口
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
func TaobaoJstInteractivePointDecrease(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointDecreaseAPIRequest, resp *jstinteractive.TaobaoJstInteractivePointDecreaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
