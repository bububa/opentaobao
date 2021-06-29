package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-店铺关联推荐 APIResponse
taobao.tbk.shop.recommend.get

入参卖家id，可推荐与此店铺相关联的相关店铺。
*/
type TaobaoTbkShopRecommendGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkShopRecommendGetResponse
}

type TaobaoTbkShopRecommendGetResponse struct {
    XMLName xml.Name `xml:"tbk_shop_recommend_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 淘宝客店铺
    
    Results   []NTbkShop `json:"results,omitempty" xml:"results>n_tbk_shop,omitempty"`
    
    
}
