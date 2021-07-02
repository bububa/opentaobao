package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemBaseAdd 【API3.0】度假线路商品发布接口
// taobao.alitrip.travel.item.base.add
//
// 旅行度假新商品发布接口。目前支持的类目包括：境内跟团游、出境跟团游、境内自由行、出境自由行、境内当地玩乐、境外玩乐套餐、境内邮轮、国际邮轮
func TaobaoAlitripTravelItemBaseAdd(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemBaseAddAPIRequest, session string) (*travel.TaobaoAlitripTravelItemBaseAddAPIResponse, error) {
	var resp travel.TaobaoAlitripTravelItemBaseAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
