package perfect

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaPerfectPerformanceLocalitemPublish 同城购定制化发品
// alibaba.perfect.performance.localitem.publish
//
// 同城购业务定制化发品接口，同城购业务线专用
func AlibabaPerfectPerformanceLocalitemPublish(ctx context.Context, clt *core.SDKClient, req *perfect.AlibabaPerfectPerformanceLocalitemPublishAPIRequest, resp *perfect.AlibabaPerfectPerformanceLocalitemPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
