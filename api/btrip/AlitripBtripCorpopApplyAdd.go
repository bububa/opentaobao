package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyAdd 【商旅】isv添加审批单
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
func AlitripBtripCorpopApplyAdd(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyAddAPIRequest, resp *btrip.AlitripBtripCorpopApplyAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
