package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapipoiget 获取景点（POI）信息
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
func Taobaoalitriptravelfscrouteapipoiget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapipoigetAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapipoigetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapipoigetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
