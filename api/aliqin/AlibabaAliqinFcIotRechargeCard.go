package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotRechargeCard 按终端号订购增值业务
// alibaba.aliqin.fc.iot.rechargeCard
//
// 按终端号订购增值业务
func AlibabaAliqinFcIotRechargeCard(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotRechargeCardAPIRequest, resp *aliqin.AlibabaAliqinFcIotRechargeCardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
