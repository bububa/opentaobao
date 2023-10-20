package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinpoisearch 景点poi搜索-阿信
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
func Taobaoalitriptravelaxinpoisearch(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinpoisearchAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinpoisearchAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinpoisearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
