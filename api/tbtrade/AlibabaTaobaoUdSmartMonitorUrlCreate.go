package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// AlibabaTaobaoUdSmartMonitorUrlCreate UD效果外投投放监测链接生成
// alibaba.taobao.ud.smart.monitor.url.create
//
// UD效果外投投放监测链接生成
func AlibabaTaobaoUdSmartMonitorUrlCreate(ctx context.Context, clt *core.SDKClient, req *tbtrade.AlibabaTaobaoUdSmartMonitorUrlCreateAPIRequest, resp *tbtrade.AlibabaTaobaoUdSmartMonitorUrlCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
