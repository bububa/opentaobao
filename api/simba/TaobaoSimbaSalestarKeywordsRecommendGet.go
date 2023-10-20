package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestarkeywordsrecommendget 销量明星api相关接口
// taobao.simba.salestar.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
func Taobaosimbasalestarkeywordsrecommendget(clt *core.SDKClient, req *simba.TaobaosimbasalestarkeywordsrecommendgetAPIRequest, session string) (*simba.TaobaosimbasalestarkeywordsrecommendgetAPIResponse, error) {
	var resp simba.TaobaosimbasalestarkeywordsrecommendgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
