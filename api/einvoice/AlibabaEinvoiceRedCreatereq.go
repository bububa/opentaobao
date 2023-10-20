package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceredcreatereq 发票冲红接口
// alibaba.einvoice.red.createreq
//
// 发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
func Alibabaeinvoiceredcreatereq(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceredcreatereqAPIRequest, session string) (*einvoice.AlibabaeinvoiceredcreatereqAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceredcreatereqAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
