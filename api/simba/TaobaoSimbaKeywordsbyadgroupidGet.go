package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsbyadgroupidGet 取得一个推广组的所有关键词
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
func TaobaoSimbaKeywordsbyadgroupidGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsbyadgroupidGetAPIRequest, resp *simba.TaobaoSimbaKeywordsbyadgroupidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
