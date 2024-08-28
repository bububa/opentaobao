package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointCollectionstandpointQuery 查询收藏口径
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
func AlibabaLegalStandpointCollectionstandpointQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointCollectionstandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointCollectionstandpointQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
