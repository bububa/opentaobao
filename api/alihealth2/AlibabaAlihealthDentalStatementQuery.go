package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalStatementQuery ISV查询对账单
// alibaba.alihealth.dental.statement.query
//
// ISV查询对账单
func AlibabaAlihealthDentalStatementQuery(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStatementQueryAPIRequest, resp *alihealth2.AlibabaAlihealthDentalStatementQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
