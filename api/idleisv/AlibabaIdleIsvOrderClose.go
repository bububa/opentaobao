package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvorderclose 服务商闲鱼卖家主动关闭订单
// alibaba.idle.isv.order.close
//
// 供外部服务商 isv 提供卖家主动关闭交易订单的功能
func Alibabaidleisvorderclose(clt *core.SDKClient, req *idleisv.AlibabaidleisvordercloseAPIRequest, session string) (*idleisv.AlibabaidleisvordercloseAPIResponse, error) {
	var resp idleisv.AlibabaidleisvordercloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
