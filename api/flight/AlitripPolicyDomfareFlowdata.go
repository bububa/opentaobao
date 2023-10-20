package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicydomfareflowdata 比价工具流量详情
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
func Alitrippolicydomfareflowdata(clt *core.SDKClient, req *flight.AlitrippolicydomfareflowdataAPIRequest, session string) (*flight.AlitrippolicydomfareflowdataAPIResponse, error) {
	var resp flight.AlitrippolicydomfareflowdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
