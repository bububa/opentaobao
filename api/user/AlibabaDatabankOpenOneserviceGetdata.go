package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaDatabankOpenOneserviceGetdata 瓴羊DaaS消费者运营CGP取数接口
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
func AlibabaDatabankOpenOneserviceGetdata(ctx context.Context, clt *core.SDKClient, req *user.AlibabaDatabankOpenOneserviceGetdataAPIRequest, resp *user.AlibabaDatabankOpenOneserviceGetdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
