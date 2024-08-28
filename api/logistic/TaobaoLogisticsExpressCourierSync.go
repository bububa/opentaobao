package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressCourierSync 快递公司同步小件员信息
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
func TaobaoLogisticsExpressCourierSync(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressCourierSyncAPIRequest, resp *logistic.TaobaoLogisticsExpressCourierSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
