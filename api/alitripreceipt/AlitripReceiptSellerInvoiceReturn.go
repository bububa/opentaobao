package alitripreceipt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// Alitripreceiptsellerinvoicereturn 飞猪发票商家回调接口
// alitrip.receipt.seller.invoice.return
//
// 飞猪发票回调接口
func Alitripreceiptsellerinvoicereturn(clt *core.SDKClient, req *alitripreceipt.AlitripreceiptsellerinvoicereturnAPIRequest, session string) (*alitripreceipt.AlitripreceiptsellerinvoicereturnAPIResponse, error) {
	var resp alitripreceipt.AlitripreceiptsellerinvoicereturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
