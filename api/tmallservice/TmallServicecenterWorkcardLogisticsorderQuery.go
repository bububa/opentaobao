package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardLogisticsorderQuery 物流单信息查询
// tmall.servicecenter.workcard.logisticsorder.query
//
// 物流订单信息查询API
func TmallServicecenterWorkcardLogisticsorderQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardLogisticsorderQueryAPIRequest, resp *tmallservice.TmallServicecenterWorkcardLogisticsorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
