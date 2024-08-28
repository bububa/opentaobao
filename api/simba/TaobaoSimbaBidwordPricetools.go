package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaBidwordPricetools 关键词出价指导工具（新）
// taobao.simba.bidword.pricetools
//
// 关键词出价指导工具（新）
func TaobaoSimbaBidwordPricetools(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaBidwordPricetoolsAPIRequest, resp *simba.TaobaoSimbaBidwordPricetoolsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
