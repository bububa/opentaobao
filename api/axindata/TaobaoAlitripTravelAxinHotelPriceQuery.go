package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelpricequery 阿信酒店分销-实时报价查询
// taobao.alitrip.travel.axin.hotel.price.query
//
// 酒店实时报价查询
func Taobaoalitriptravelaxinhotelpricequery(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelpricequeryAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelpricequeryAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelpricequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
