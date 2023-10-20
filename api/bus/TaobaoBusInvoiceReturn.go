package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusinvoicereturn 发票回调接口
// taobao.bus.invoice.return
//
// 汽车票发票回调接口
func Taobaobusinvoicereturn(clt *core.SDKClient, req *bus.TaobaobusinvoicereturnAPIRequest, session string) (*bus.TaobaobusinvoicereturnAPIResponse, error) {
	var resp bus.TaobaobusinvoicereturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
