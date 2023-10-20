package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsrecommendget 取得一个推广组的推荐关键词列表
// taobao.simba.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
func Taobaosimbakeywordsrecommendget(clt *core.SDKClient, req *simba.TaobaosimbakeywordsrecommendgetAPIRequest, session string) (*simba.TaobaosimbakeywordsrecommendgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsrecommendgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
