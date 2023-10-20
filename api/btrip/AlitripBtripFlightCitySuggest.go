package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightCitySuggest 机票城市搜索
// alitrip.btrip.flight.city.suggest
//
// 提供机票城市搜索接口，提高OA用户对接效率
func AlitripBtripFlightCitySuggest(clt *core.SDKClient, req *btrip.AlitripBtripFlightCitySuggestAPIRequest, resp *btrip.AlitripBtripFlightCitySuggestAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
