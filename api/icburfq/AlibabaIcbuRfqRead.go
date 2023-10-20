package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicburfqread 是否已读RFQ
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
func Alibabaicburfqread(clt *core.SDKClient, req *icburfq.AlibabaicburfqreadAPIRequest, session string) (*icburfq.AlibabaicburfqreadAPIResponse, error) {
	var resp icburfq.AlibabaicburfqreadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
