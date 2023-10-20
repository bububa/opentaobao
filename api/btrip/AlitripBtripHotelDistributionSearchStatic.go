package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionsearchstatic 商旅酒店api分销-酒店静态信息接口
// alitrip.btrip.hotel.distribution.search.static
//
// 商旅酒店api分销-酒店静态信息接口
func Alitripbtriphoteldistributionsearchstatic(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionsearchstaticAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionsearchstaticAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionsearchstaticAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
