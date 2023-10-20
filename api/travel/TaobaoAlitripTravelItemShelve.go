package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemShelve 【API3.0】度假线路商品上下架接口
// taobao.alitrip.travel.item.shelve
//
// 旅行度假新商品发布接口 第三版：度假商品上下架接口
// 注意：定时上下架功能，目前只支持接送、租车类目。
func TaobaoAlitripTravelItemShelve(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemShelveAPIRequest, session string) (*travel.TaobaoAlitripTravelItemShelveAPIResponse, error) {
	var resp travel.TaobaoAlitripTravelItemShelveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
