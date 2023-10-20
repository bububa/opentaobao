package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicydomfarecompare 比价工具
// alitrip.policy.domfare.compare
//
// 比价工具
func Alitrippolicydomfarecompare(clt *core.SDKClient, req *flight.AlitrippolicydomfarecompareAPIRequest, session string) (*flight.AlitrippolicydomfarecompareAPIResponse, error) {
	var resp flight.AlitrippolicydomfarecompareAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
