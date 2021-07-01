package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

/* AlibabaEinvoiceBillSync
结算单同步
alibaba.einvoice.bill.sync

电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票 */
func AlibabaEinvoiceBillSync(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceBillSyncAPIRequest, session string) (*einvoice.AlibabaEinvoiceBillSyncAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceBillSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
