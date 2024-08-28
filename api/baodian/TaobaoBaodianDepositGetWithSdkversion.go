package baodian

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoBaodianDepositGetWithSdkversion 查询用户宝点信息（带sdk版本，已迁移）
// taobao.baodian.deposit.get.with.sdkversion
//
// 获取用户宝点信息（带sdk版本，已迁移）
func TaobaoBaodianDepositGetWithSdkversion(ctx context.Context, clt *core.SDKClient, req *baodian.TaobaoBaodianDepositGetWithSdkversionAPIRequest, resp *baodian.TaobaoBaodianDepositGetWithSdkversionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
