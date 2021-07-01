package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

/* AlibabaIcbuRfqSearch
查询RFQ
alibaba.icbu.rfq.search

用于查询RFQ的信息 */
func AlibabaIcbuRfqSearch(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqSearchAPIRequest, session string) (*icburfq.AlibabaIcbuRfqSearchAPIResponse, error) {
	var resp icburfq.AlibabaIcbuRfqSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
