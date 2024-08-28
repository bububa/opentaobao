package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightBasicDataCityQueryAll 机票基础数据城市数据查询
// alitrip.flight.basic.data.city.queryAll
//
// 机票基础数据城市数据查询top接口
func AlitripFlightBasicDataCityQueryAll(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightBasicDataCityQueryAllAPIRequest, resp *flightuppc.AlitripFlightBasicDataCityQueryAllAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
