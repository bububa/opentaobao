package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusLogisticsPackagesNotice 物流多包裹通知
// taobao.rdc.aligenius.logistics.packages.notice
//
// 订单发货之后，如果订单拆包、补发、赠品等场景，需要将多余包裹信息触达消费者, 大促会降级
func TaobaoRdcAligeniusLogisticsPackagesNotice(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest, resp *logistic.TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
