package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCreativeUpdate 销量明星更新创意相关接口
// taobao.simba.salestar.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
func TaobaoSimbaSalestarCreativeUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativeUpdateAPIRequest, resp *simba.TaobaoSimbaSalestarCreativeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
