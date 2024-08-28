package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpWordpackageFindlist 词包列表查询
// taobao.universalbp.wordpackage.findlist
//
// 根据计划+单元id，查绑定的词包列表
func TaobaoUniversalbpWordpackageFindlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpWordpackageFindlistAPIRequest, resp *simba.TaobaoUniversalbpWordpackageFindlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
