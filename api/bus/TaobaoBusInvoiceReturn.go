package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusInvoiceReturn 发票回调接口
// taobao.bus.invoice.return
//
// 汽车票发票回调接口
func TaobaoBusInvoiceReturn(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusInvoiceReturnAPIRequest, resp *bus.TaobaoBusInvoiceReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
