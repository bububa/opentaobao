package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinpoidetailquery 景点poi详情查询-阿信
// taobao.alitrip.travel.axin.poi.detail.query
//
// 景点poi详情查询-阿信
func Taobaoalitriptravelaxinpoidetailquery(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinpoidetailqueryAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinpoidetailqueryAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinpoidetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
