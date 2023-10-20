package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqSearch 查询RFQ
// alibaba.icbu.rfq.search
//
// 用于查询RFQ的信息
func AlibabaIcbuRfqSearch(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqSearchAPIRequest, resp *icburfq.AlibabaIcbuRfqSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
