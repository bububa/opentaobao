package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsHluserGet 订单全链路用户信息获取
// taobao.jds.hluser.get
//
// 订单全链路用户信息获取
func TaobaoJdsHluserGet(clt *core.SDKClient, req *jst.TaobaoJdsHluserGetAPIRequest, resp *jst.TaobaoJdsHluserGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
