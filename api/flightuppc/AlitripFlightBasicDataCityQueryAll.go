package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightBasicDataCityQueryAll 机票基础数据城市数据查询
// alitrip.flight.basic.data.city.queryAll
//
// 机票基础数据城市数据查询top接口
func AlitripFlightBasicDataCityQueryAll(clt *core.SDKClient, req *flightuppc.AlitripFlightBasicDataCityQueryAllAPIRequest, resp *flightuppc.AlitripFlightBasicDataCityQueryAllAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
