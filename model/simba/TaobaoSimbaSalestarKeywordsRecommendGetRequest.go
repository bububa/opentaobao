package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星api相关接口 APIRequest
taobao.simba.salestar.keywords.recommend.get

取得一个推广组的推荐关键词列表
*/
type TaobaoSimbaSalestarKeywordsRecommendGetRequest struct {
    model.Params

    // 推广组ID
    adgroupId   int64 

    // 产品类型101001005代表标准推广，101001014代表销量明星
    productId   int64 

}

func NewTaobaoSimbaSalestarKeywordsRecommendGetRequest() *TaobaoSimbaSalestarKeywordsRecommendGetRequest{
    return &TaobaoSimbaSalestarKeywordsRecommendGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.keywords.recommend.get"
}

func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarKeywordsRecommendGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaSalestarKeywordsRecommendGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetProductId() int64 {
    return r.productId
}

