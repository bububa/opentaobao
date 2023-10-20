package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCreativeAdd （新）新建创意
// taobao.simba.salestar.creative.add
//
// 创建一个创意
func TaobaoSimbaSalestarCreativeAdd(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativeAddAPIRequest, resp *simba.TaobaoSimbaSalestarCreativeAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
