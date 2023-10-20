package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceamountcheck 开票量核对接口
// alibaba.einvoice.amount.check
//
// 跟开票服务商核对历史开票量，用来对账
func Alibabaeinvoiceamountcheck(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceamountcheckAPIRequest, session string) (*einvoice.AlibabaeinvoiceamountcheckAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceamountcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
