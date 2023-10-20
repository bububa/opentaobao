package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightcitysuggest 机票城市搜索
// alitrip.btrip.flight.city.suggest
//
// 提供机票城市搜索接口，提高OA用户对接效率
func Alitripbtripflightcitysuggest(clt *core.SDKClient, req *btrip.AlitripbtripflightcitysuggestAPIRequest, session string) (*btrip.AlitripbtripflightcitysuggestAPIResponse, error) {
	var resp btrip.AlitripbtripflightcitysuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
