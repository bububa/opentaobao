package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemBaseModify 【API3.0】度假线路商品编辑接口
// taobao.alitrip.travel.item.base.modify
//
// 旅行度假新商品基本信息修改接口 第三版。提供商家通过TOP API方式修改商品除sku外的基本信息。
func TaobaoAlitripTravelItemBaseModify(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemBaseModifyAPIRequest, session string) (*travel.TaobaoAlitripTravelItemBaseModifyAPIResponse, error) {
	var resp travel.TaobaoAlitripTravelItemBaseModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
