package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaCreativeDelete
删除创意
taobao.simba.creative.delete

删除一个创意 */
func TaobaoSimbaCreativeDelete(clt *core.SDKClient, req *simba.TaobaoSimbaCreativeDeleteAPIRequest, session string) (*simba.TaobaoSimbaCreativeDeleteAPIResponse, error) {
	var resp simba.TaobaoSimbaCreativeDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
