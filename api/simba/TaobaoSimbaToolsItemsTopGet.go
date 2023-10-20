package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaToolsItemsTopGet 取得一个关键词的推广组排名列表
// taobao.simba.tools.items.top.get
//
// 取得一个关键词的推广组排名列表
func TaobaoSimbaToolsItemsTopGet(clt *core.SDKClient, req *simba.TaobaoSimbaToolsItemsTopGetAPIRequest, resp *simba.TaobaoSimbaToolsItemsTopGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
