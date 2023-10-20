package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicburfqsearch 查询RFQ
// alibaba.icbu.rfq.search
//
// 用于查询RFQ的信息
func Alibabaicburfqsearch(clt *core.SDKClient, req *icburfq.AlibabaicburfqsearchAPIRequest, session string) (*icburfq.AlibabaicburfqsearchAPIResponse, error) {
	var resp icburfq.AlibabaicburfqsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
