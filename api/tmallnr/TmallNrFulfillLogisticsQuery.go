package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillLogisticsQuery 定时送和极速达配送物流信息查询
// tmall.nr.fulfill.logistics.query
//
// 发布定时送&amp;极速达物流信息查询服务
func TmallNrFulfillLogisticsQuery(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrFulfillLogisticsQueryAPIRequest, resp *tmallnr.TmallNrFulfillLogisticsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
