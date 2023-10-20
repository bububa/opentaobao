package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhoteldetailquery 阿信酒店分销-标准酒店详情查询
// taobao.alitrip.travel.axin.hotel.detail.query
//
// 标准酒店详情查询
func Taobaoalitriptravelaxinhoteldetailquery(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhoteldetailqueryAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhoteldetailqueryAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhoteldetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
