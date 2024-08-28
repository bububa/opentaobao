package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyGet 【商旅】查询审批单
// alitrip.btrip.corpop.apply.get
//
// 【商旅】查询审批单
func AlitripBtripCorpopApplyGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyGetAPIRequest, resp *btrip.AlitripBtripCorpopApplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
