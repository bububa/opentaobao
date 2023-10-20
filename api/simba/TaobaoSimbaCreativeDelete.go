package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCreativeDelete 删除创意
// taobao.simba.creative.delete
//
// 删除一个创意
func TaobaoSimbaCreativeDelete(clt *core.SDKClient, req *simba.TaobaoSimbaCreativeDeleteAPIRequest, resp *simba.TaobaoSimbaCreativeDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
