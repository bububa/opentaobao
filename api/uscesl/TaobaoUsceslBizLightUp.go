package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizLightUp 价签LED等点亮
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
func TaobaoUsceslBizLightUp(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizLightUpAPIRequest, resp *uscesl.TaobaoUsceslBizLightUpAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
