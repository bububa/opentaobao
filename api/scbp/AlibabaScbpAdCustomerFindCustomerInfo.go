package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCustomerFindCustomerInfo 查询客户信息
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
func AlibabaScbpAdCustomerFindCustomerInfo(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCustomerFindCustomerInfoAPIRequest, resp *scbp.AlibabaScbpAdCustomerFindCustomerInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
