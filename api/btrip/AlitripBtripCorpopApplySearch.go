package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplySearch 【商旅】搜索审批单列表
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
func AlitripBtripCorpopApplySearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplySearchAPIRequest, resp *btrip.AlitripBtripCorpopApplySearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
