package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsPricevonSet 设置一批关键词的信息
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
func TaobaoSimbaKeywordsPricevonSet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsPricevonSetAPIRequest, session string) (*simba.TaobaoSimbaKeywordsPricevonSetAPIResponse, error) {
	var resp simba.TaobaoSimbaKeywordsPricevonSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
