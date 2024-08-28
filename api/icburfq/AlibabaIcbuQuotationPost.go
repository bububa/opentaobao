package icburfq

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuQuotationPost 供应商提交报价
// alibaba.icbu.quotation.post
//
// 供应商对RFQ进行报价
func AlibabaIcbuQuotationPost(ctx context.Context, clt *core.SDKClient, req *icburfq.AlibabaIcbuQuotationPostAPIRequest, resp *icburfq.AlibabaIcbuQuotationPostAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
