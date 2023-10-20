package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceRedCreatereq 发票冲红接口
// alibaba.einvoice.red.createreq
//
// 发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
func AlibabaEinvoiceRedCreatereq(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceRedCreatereqAPIRequest, resp *einvoice.AlibabaEinvoiceRedCreatereqAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
