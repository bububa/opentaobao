package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCreativesChangedGet 分页获取修改过的广告创意ID和修改时间
// taobao.simba.creatives.changed.get
//
// 分页获取修改过的广告创意ID和修改时间
func TaobaoSimbaCreativesChangedGet(clt *core.SDKClient, req *simba.TaobaoSimbaCreativesChangedGetAPIRequest, resp *simba.TaobaoSimbaCreativesChangedGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
