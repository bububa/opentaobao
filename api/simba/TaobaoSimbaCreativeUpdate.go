package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaCreativeUpdate
修改创意与
taobao.simba.creative.update

更新一个创意的信息，可以设置创意标题、创意图片 */
func TaobaoSimbaCreativeUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCreativeUpdateAPIRequest, session string) (*simba.TaobaoSimbaCreativeUpdateAPIResponse, error) {
	var resp simba.TaobaoSimbaCreativeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
