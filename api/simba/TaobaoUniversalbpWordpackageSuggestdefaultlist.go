package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpWordpackageSuggestdefaultlist 建议默认关键词包
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
func TaobaoUniversalbpWordpackageSuggestdefaultlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest, resp *simba.TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
