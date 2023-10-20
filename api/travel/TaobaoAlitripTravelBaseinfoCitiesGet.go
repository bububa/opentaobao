package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Taobaoalitriptravelbaseinfocitiesget 【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询
// taobao.alitrip.travel.baseinfo.cities.get
//
// 旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。
func Taobaoalitriptravelbaseinfocitiesget(clt *core.SDKClient, req *travel.TaobaoalitriptravelbaseinfocitiesgetAPIRequest, session string) (*travel.TaobaoalitriptravelbaseinfocitiesgetAPIResponse, error) {
	var resp travel.TaobaoalitriptravelbaseinfocitiesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
