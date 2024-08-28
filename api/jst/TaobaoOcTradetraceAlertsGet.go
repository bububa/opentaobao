package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcTradetraceAlertsGet 异常订单信息获取
// taobao.oc.tradetrace.alerts.get
//
// 提供订单预警模块的数据查询接口
func TaobaoOcTradetraceAlertsGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoOcTradetraceAlertsGetAPIRequest, resp *jst.TaobaoOcTradetraceAlertsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
