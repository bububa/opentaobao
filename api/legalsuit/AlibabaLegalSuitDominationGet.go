package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitDominationGet 查询管辖信息
// alibaba.legal.suit.domination.get
//
// 查询管辖信息
func AlibabaLegalSuitDominationGet(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitDominationGetAPIRequest, resp *legalsuit.AlibabaLegalSuitDominationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
