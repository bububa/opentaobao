package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletradecarperform 二手车寄卖履约接口
// alibaba.idle.trade.car.perform
//
// 二手车寄卖履约接口
func Alibabaidletradecarperform(clt *core.SDKClient, req *idle.AlibabaidletradecarperformAPIRequest, session string) (*idle.AlibabaidletradecarperformAPIResponse, error) {
	var resp idle.AlibabaidletradecarperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
