package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarKeywordsRecommendGet 销量明星api相关接口
// taobao.simba.salestar.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
func TaobaoSimbaSalestarKeywordsRecommendGet(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest, session string) (*simba.TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse, error) {
	var resp simba.TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
