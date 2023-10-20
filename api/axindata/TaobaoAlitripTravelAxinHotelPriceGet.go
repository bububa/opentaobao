package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelpriceget 酒店报价服务-阿信
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
func Taobaoalitriptravelaxinhotelpriceget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelpricegetAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelpricegetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelpricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
