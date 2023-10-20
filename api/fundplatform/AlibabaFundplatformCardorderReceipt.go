package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardorderReceipt 通知确认收货
// alibaba.fundplatform.cardorder.receipt
//
// 告知卡商这一批储值卡已经被用户确认收货
func AlibabaFundplatformCardorderReceipt(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderReceiptAPIRequest, resp *fundplatform.AlibabaFundplatformCardorderReceiptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
