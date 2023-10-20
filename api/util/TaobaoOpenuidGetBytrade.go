package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoOpenuidGetBytrade 通过订单获取对应买家的openUID
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
func TaobaoOpenuidGetBytrade(clt *core.SDKClient, req *util.TaobaoOpenuidGetBytradeAPIRequest, resp *util.TaobaoOpenuidGetBytradeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
