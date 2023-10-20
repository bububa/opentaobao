package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqRead 是否已读RFQ
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
func AlibabaIcbuRfqRead(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqReadAPIRequest, resp *icburfq.AlibabaIcbuRfqReadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
