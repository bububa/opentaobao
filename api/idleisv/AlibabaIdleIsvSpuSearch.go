package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvSpuSearch spu搜索接口
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
func AlibabaIdleIsvSpuSearch(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvSpuSearchAPIRequest, resp *idleisv.AlibabaIdleIsvSpuSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
