package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Taobaoalitriptravelbaseinfocruiseget 【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取
// taobao.alitrip.travel.baseinfo.cruise.get
//
// 旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
func Taobaoalitriptravelbaseinfocruiseget(clt *core.SDKClient, req *travel.TaobaoalitriptravelbaseinfocruisegetAPIRequest, session string) (*travel.TaobaoalitriptravelbaseinfocruisegetAPIResponse, error) {
	var resp travel.TaobaoalitriptravelbaseinfocruisegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
