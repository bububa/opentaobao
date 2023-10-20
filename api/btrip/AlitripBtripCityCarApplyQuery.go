package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCityCarApplyQuery 三方市内用车申请单查询
// alitrip.btrip.city.car.apply.query
//
// 三方市内用车申请单查询
func AlitripBtripCityCarApplyQuery(clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyQueryAPIRequest, resp *btrip.AlitripBtripCityCarApplyQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
