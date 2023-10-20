package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleisvgoosefishordercreate 闲鱼三方安康容器订单创建
// alibaba.idle.isv.goosefish.order.create
//
// 闲鱼三方安康容器订单创建
func Alibabaidleisvgoosefishordercreate(clt *core.SDKClient, req *idle.AlibabaidleisvgoosefishordercreateAPIRequest, session string) (*idle.AlibabaidleisvgoosefishordercreateAPIResponse, error) {
	var resp idle.AlibabaidleisvgoosefishordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
