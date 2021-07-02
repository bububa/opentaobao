package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordFindbyids （新）根据一堆关键词ids获取关键词
// taobao.simba.keyword.findbyids
//
// 根据一个关键词Id列表取得一组关键词
func TaobaoSimbaKeywordFindbyids(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordFindbyidsAPIRequest, session string) (*simba.TaobaoSimbaKeywordFindbyidsAPIResponse, error) {
	var resp simba.TaobaoSimbaKeywordFindbyidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
