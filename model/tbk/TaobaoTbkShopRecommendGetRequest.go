package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-店铺关联推荐 APIRequest
taobao.tbk.shop.recommend.get

入参卖家id，可推荐与此店铺相关联的相关店铺。
*/
type TaobaoTbkShopRecommendGetRequest struct {
    model.Params

    // 需返回的字段列表
    fields   string 

    // 卖家Id
    userId   int64 

    // 返回数量，默认20，最大值40
    count   int64 

    // 链接形式：1：PC，2：无线，默认：１
    platform   int64 

}

func NewTaobaoTbkShopRecommendGetRequest() *TaobaoTbkShopRecommendGetRequest{
    return &TaobaoTbkShopRecommendGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkShopRecommendGetRequest) GetApiMethodName() string {
    return "taobao.tbk.shop.recommend.get"
}

func (r TaobaoTbkShopRecommendGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkShopRecommendGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTbkShopRecommendGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoTbkShopRecommendGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoTbkShopRecommendGetRequest) GetUserId() int64 {
    return r.userId
}

func (r *TaobaoTbkShopRecommendGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r TaobaoTbkShopRecommendGetRequest) GetCount() int64 {
    return r.count
}

func (r *TaobaoTbkShopRecommendGetRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoTbkShopRecommendGetRequest) GetPlatform() int64 {
    return r.platform
}

