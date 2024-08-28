package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtItemMainSynchronize 家装新零售主商品同步至阿里
// tmall.nrt.item.main.synchronize
//
// 同步卖场存量线下商品到阿里
func TmallNrtItemMainSynchronize(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtItemMainSynchronizeAPIRequest, resp *nrt.TmallNrtItemMainSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
