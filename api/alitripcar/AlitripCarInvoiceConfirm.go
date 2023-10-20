package alitripcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripcar"
)

// AlitripCarInvoiceConfirm 发票确认接口
// alitrip.car.invoice.confirm
//
// 飞猪发票回调接口
func AlitripCarInvoiceConfirm(clt *core.SDKClient, req *alitripcar.AlitripCarInvoiceConfirmAPIRequest, resp *alitripcar.AlitripCarInvoiceConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
