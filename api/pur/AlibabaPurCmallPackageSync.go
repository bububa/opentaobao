package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCmallPackageSync 套餐同步
// alibaba.pur.cmall.package.sync
//
// 套餐同步
func AlibabaPurCmallPackageSync(clt *core.SDKClient, req *pur.AlibabaPurCmallPackageSyncAPIRequest, resp *pur.AlibabaPurCmallPackageSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
