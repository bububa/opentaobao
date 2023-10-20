package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeDeviceReturn 服务商回传客户端设备列表
// alibaba.einvoice.income.device.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
func AlibabaEinvoiceIncomeDeviceReturn(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeDeviceReturnAPIRequest, session string) (*einvoice.AlibabaEinvoiceIncomeDeviceReturnAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceIncomeDeviceReturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
