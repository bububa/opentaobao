package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceDeviceOrderQuery 查询税控设备加盘订购单详情
// alibaba.einvoice.device.order.query
//
// 查询税控设备订购单详情
func AlibabaEinvoiceDeviceOrderQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceDeviceOrderQueryAPIRequest, resp *einvoice.AlibabaEinvoiceDeviceOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
