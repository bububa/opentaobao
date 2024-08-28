package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCmallPackageSync 套餐同步
// alibaba.pur.cmall.package.sync
//
// 套餐同步
func AlibabaPurCmallPackageSync(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurCmallPackageSyncAPIRequest, resp *pur.AlibabaPurCmallPackageSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
