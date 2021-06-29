package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星api相关接口 API请求
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

// 初始化TaobaoSimbaSalestarKeywordsRecommendGetRequest对象
func NewTaobaoSimbaSalestarKeywordsRecommendGetRequest() *TaobaoSimbaSalestarKeywordsRecommendGetRequest{
    return &TaobaoSimbaSalestarKeywordsRecommendGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.keywords.recommend.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaSalestarKeywordsRecommendGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// ProductId Setter
// 产品类型101001005代表标准推广，101001014代表销量明星
func (r *TaobaoSimbaSalestarKeywordsRecommendGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoSimbaSalestarKeywordsRecommendGetRequest) GetProductId() int64 {
    return r.productId
}
