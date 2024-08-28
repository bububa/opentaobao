package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeQuery 查询公益3小时公益时汇总
// alibaba.charity.charitytime.query
//
// 查询公益3小时公益时汇总
func AlibabaCharityCharitytimeQuery(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeQueryAPIRequest, resp *charity.AlibabaCharityCharitytimeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
