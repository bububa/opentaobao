package perfect

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaPerfectPerformanceItemQuery 商品完美履约信息查询
// alibaba.perfect.performance.item.query
//
// 同城零售商品完美履约信息查询
func AlibabaPerfectPerformanceItemQuery(ctx context.Context, clt *core.SDKClient, req *perfect.AlibabaPerfectPerformanceItemQueryAPIRequest, resp *perfect.AlibabaPerfectPerformanceItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
