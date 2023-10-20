package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripflightbasicdatacityqueryAll 机票基础数据城市数据查询
// alitrip.flight.basic.data.city.queryAll
//
// 机票基础数据城市数据查询top接口
func AlitripflightbasicdatacityqueryAll(clt *core.SDKClient, req *flightuppc.AlitripflightbasicdatacityqueryAllAPIRequest, session string) (*flightuppc.AlitripflightbasicdatacityqueryAllAPIResponse, error) {
	var resp flightuppc.AlitripflightbasicdatacityqueryAllAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
