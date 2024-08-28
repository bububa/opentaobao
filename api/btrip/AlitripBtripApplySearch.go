package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApplySearch 搜索审批单
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
func AlitripBtripApplySearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripApplySearchAPIRequest, resp *btrip.AlitripBtripApplySearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
