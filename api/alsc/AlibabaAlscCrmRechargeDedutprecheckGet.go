package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeDedutprecheckGet 储值核销预先校验
// alibaba.alsc.crm.recharge.dedutprecheck.get
//
// 储值核销预先校验接口
func AlibabaAlscCrmRechargeDedutprecheckGet(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest, resp *alsc.AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
