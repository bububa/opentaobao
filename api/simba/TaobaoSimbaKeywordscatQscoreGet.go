package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordscatQscoreGet 取得一个推广组的所有关键词和类目出价的质量得分
// taobao.simba.keywordscat.qscore.get
//
// 取得一个推广组的所有关键词和类目出价的质量得分列表
func TaobaoSimbaKeywordscatQscoreGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordscatQscoreGetAPIRequest, resp *simba.TaobaoSimbaKeywordscatQscoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
