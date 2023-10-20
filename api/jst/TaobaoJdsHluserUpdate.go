package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsHluserUpdate 订单全链路用户信息修改
// taobao.jds.hluser.update
//
// 订单全链路用户信息修改，比如是否开放买家端展示
func TaobaoJdsHluserUpdate(clt *core.SDKClient, req *jst.TaobaoJdsHluserUpdateAPIRequest, resp *jst.TaobaoJdsHluserUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
