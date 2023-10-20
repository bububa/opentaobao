package alitripreceipt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// Alitripreceiptsellerinvoicered 飞猪发票冲红
// alitrip.receipt.seller.invoice.red
//
// 飞猪发票创建
func Alitripreceiptsellerinvoicered(clt *core.SDKClient, req *alitripreceipt.AlitripreceiptsellerinvoiceredAPIRequest, session string) (*alitripreceipt.AlitripreceiptsellerinvoiceredAPIResponse, error) {
	var resp alitripreceipt.AlitripreceiptsellerinvoiceredAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
