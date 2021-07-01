package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaToolsItemsTopGet
取得一个关键词的推广组排名列表
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表 */
func TaobaoSimbaToolsItemsTopGet(clt *core.SDKClient, req *simba.TaobaoSimbaToolsItemsTopGetAPIRequest, session string) (*simba.TaobaoSimbaToolsItemsTopGetAPIResponse, error) {
	var resp simba.TaobaoSimbaToolsItemsTopGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
