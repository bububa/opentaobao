package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-店铺关联推荐 APIResponse
taobao.tbk.shop.recommend.get

入参卖家id，可推荐与此店铺相关联的相关店铺。
*/
type TaobaoTbkShopRecommendGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkShopRecommendGetResponse `json:"taobao_tbk_shop_recommend_get_response,omitempty"`
}

type TaobaoTbkShopRecommendGetResponse struct {

    // 淘宝客店铺
    Results   []NTbkShop `json:"results,omitempty"`

}
