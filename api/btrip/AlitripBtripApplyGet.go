package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApplyGet 获取单个审批单
// alitrip.btrip.apply.get
//
// 获取单个审批单的详情数据
func AlitripBtripApplyGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripApplyGetAPIRequest, resp *btrip.AlitripBtripApplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
