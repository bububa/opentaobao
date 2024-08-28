package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeDeviceReturn 服务商回传客户端设备列表
// alibaba.einvoice.income.device.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
func AlibabaEinvoiceIncomeDeviceReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeDeviceReturnAPIRequest, resp *einvoice.AlibabaEinvoiceIncomeDeviceReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
