package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-店铺关联推荐 API请求
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

// 初始化TaobaoTbkShopRecommendGetRequest对象
func NewTaobaoTbkShopRecommendGetRequest() *TaobaoTbkShopRecommendGetRequest{
    return &TaobaoTbkShopRecommendGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkShopRecommendGetRequest) GetApiMethodName() string {
    return "taobao.tbk.shop.recommend.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkShopRecommendGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表
func (r *TaobaoTbkShopRecommendGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTbkShopRecommendGetRequest) GetFields() string {
    return r.fields
}
// UserId Setter
// 卖家Id
func (r *TaobaoTbkShopRecommendGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoTbkShopRecommendGetRequest) GetUserId() int64 {
    return r.userId
}
// Count Setter
// 返回数量，默认20，最大值40
func (r *TaobaoTbkShopRecommendGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r TaobaoTbkShopRecommendGetRequest) GetCount() int64 {
    return r.count
}
// Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkShopRecommendGetRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r TaobaoTbkShopRecommendGetRequest) GetPlatform() int64 {
    return r.platform
}
