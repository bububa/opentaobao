package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikePerformancePut 提交客服绩效接口
// taobao.weike.performance.put
//
// 提交客服绩效接口
func TaobaoWeikePerformancePut(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoWeikePerformancePutAPIRequest, resp *servicecenter.TaobaoWeikePerformancePutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
