package icburfq

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqSearch 查询RFQ
// alibaba.icbu.rfq.search
//
// 用于查询RFQ的信息
func AlibabaIcbuRfqSearch(ctx context.Context, clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqSearchAPIRequest, resp *icburfq.AlibabaIcbuRfqSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
