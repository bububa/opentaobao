package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaDatabankOpenOneserviceDataready 瓴羊DaaS消费者增长CGP查询DataReady
// alibaba.databank.open.oneservice.dataready
//
// 瓴羊DaaS消费者增长CGP取数接口
func AlibabaDatabankOpenOneserviceDataready(ctx context.Context, clt *core.SDKClient, req *user.AlibabaDatabankOpenOneserviceDatareadyAPIRequest, resp *user.AlibabaDatabankOpenOneserviceDatareadyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
