package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelcityget 城市列表信息查询-阿信
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
func Taobaoalitriptravelaxinhotelcityget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelcitygetAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelcitygetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelcitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
