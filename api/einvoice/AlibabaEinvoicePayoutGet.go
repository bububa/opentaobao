package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePayoutGet 获取赔付计时列表数据
// alibaba.einvoice.payout.get
//
// 获取赔付计时列表数据
func AlibabaEinvoicePayoutGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoicePayoutGetAPIRequest, session string) (*einvoice.AlibabaEinvoicePayoutGetAPIResponse, error) {
	var resp einvoice.AlibabaEinvoicePayoutGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
