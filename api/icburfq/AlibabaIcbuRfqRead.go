package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqRead 是否已读RFQ
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
func AlibabaIcbuRfqRead(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqReadAPIRequest, session string) (*icburfq.AlibabaIcbuRfqReadAPIResponse, error) {
	var resp icburfq.AlibabaIcbuRfqReadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
