package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelFastinvoiceComplete 极速开票开票请求完成
// taobao.xhotel.fastinvoice.complete
//
// 极速开票开票请求回传,用于更新航信开票请求数据
func TaobaoXhotelFastinvoiceComplete(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelFastinvoiceCompleteAPIRequest, resp *xhotelonlineorder.TaobaoXhotelFastinvoiceCompleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
