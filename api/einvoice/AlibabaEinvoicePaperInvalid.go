package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

/* AlibabaEinvoicePaperInvalid
纸票作废接口
alibaba.einvoice.paper.invalid

作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票 */
func AlibabaEinvoicePaperInvalid(clt *core.SDKClient, req *einvoice.AlibabaEinvoicePaperInvalidAPIRequest, session string) (*einvoice.AlibabaEinvoicePaperInvalidAPIResponse, error) {
	var resp einvoice.AlibabaEinvoicePaperInvalidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
