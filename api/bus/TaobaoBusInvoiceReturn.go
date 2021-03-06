package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusInvoiceReturn 发票回调接口
// taobao.bus.invoice.return
//
// 汽车票发票回调接口
func TaobaoBusInvoiceReturn(clt *core.SDKClient, req *bus.TaobaoBusInvoiceReturnAPIRequest, session string) (*bus.TaobaoBusInvoiceReturnAPIResponse, error) {
	var resp bus.TaobaoBusInvoiceReturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
