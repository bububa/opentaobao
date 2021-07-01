package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

/* TaobaoTbkShopRecommendGet
淘宝客-公用-店铺关联推荐
taobao.tbk.shop.recommend.get

入参卖家id，可推荐与此店铺相关联的相关店铺。 */
func TaobaoTbkShopRecommendGet(clt *core.SDKClient, req *tbk.TaobaoTbkShopRecommendGetAPIRequest, session string) (*tbk.TaobaoTbkShopRecommendGetAPIResponse, error) {
	var resp tbk.TaobaoTbkShopRecommendGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
