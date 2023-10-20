package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ctradeget 零售+平台查询订单
// alibaba.nlife.b2c.trade.get
//
// 查询零售+平台创建出来的订单详情
func Alibabanlifeb2ctradeget(clt *core.SDKClient, req *nlife.Alibabanlifeb2ctradegetAPIRequest, session string) (*nlife.Alibabanlifeb2ctradegetAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ctradegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
