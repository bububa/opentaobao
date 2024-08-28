package perfect

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaPerfectPerformanceLocalitemEdit 同城购定制发品编辑
// alibaba.perfect.performance.localitem.edit
//
// 同城购业务定制化发品接口，同城购业务线专用
func AlibabaPerfectPerformanceLocalitemEdit(ctx context.Context, clt *core.SDKClient, req *perfect.AlibabaPerfectPerformanceLocalitemEditAPIRequest, resp *perfect.AlibabaPerfectPerformanceLocalitemEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
