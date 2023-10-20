package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicedeviceorderquery 查询税控设备加盘订购单详情
// alibaba.einvoice.device.order.query
//
// 查询税控设备订购单详情
func Alibabaeinvoicedeviceorderquery(clt *core.SDKClient, req *einvoice.AlibabaeinvoicedeviceorderqueryAPIRequest, session string) (*einvoice.AlibabaeinvoicedeviceorderqueryAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicedeviceorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
