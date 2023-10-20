package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCreativeDelete (新)销量明星删除创意相关接口
// taobao.simba.salestar.creative.delete
//
// 删除一个创意
func TaobaoSimbaSalestarCreativeDelete(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativeDeleteAPIRequest, resp *simba.TaobaoSimbaSalestarCreativeDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
