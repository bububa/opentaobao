package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Taobaoalitriptravelitemnewquery 【API3.0】新版度假单个商品查询接口
// taobao.alitrip.travel.item.new.query
//
// 新版旅行度假新商品查询接口（单个商品查询）
func Taobaoalitriptravelitemnewquery(clt *core.SDKClient, req *travel.TaobaoalitriptravelitemnewqueryAPIRequest, session string) (*travel.TaobaoalitriptravelitemnewqueryAPIResponse, error) {
	var resp travel.TaobaoalitriptravelitemnewqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
