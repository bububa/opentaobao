package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripProjectModify 变更项目
// alitrip.btrip.project.modify
//
// 变更项目
func AlitripBtripProjectModify(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripProjectModifyAPIRequest, resp *btrip.AlitripBtripProjectModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
