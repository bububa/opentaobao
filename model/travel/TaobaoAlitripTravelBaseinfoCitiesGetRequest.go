package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 API请求
taobao.alitrip.travel.baseinfo.cities.get

旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。
*/
type TaobaoAlitripTravelBaseinfoCitiesGetRequest struct {
    model.Params
    // 1-获取目的地城市列表 2-获取出发地城市列表
    iocType   int64
    // 1-境内跟团游 2-境内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 9-境内邮轮
    catType   int64
}

// 初始化TaobaoAlitripTravelBaseinfoCitiesGetRequest对象
func NewTaobaoAlitripTravelBaseinfoCitiesGetRequest() *TaobaoAlitripTravelBaseinfoCitiesGetRequest{
    return &TaobaoAlitripTravelBaseinfoCitiesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelBaseinfoCitiesGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.baseinfo.cities.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelBaseinfoCitiesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IocType Setter
// 1-获取目的地城市列表 2-获取出发地城市列表
func (r *TaobaoAlitripTravelBaseinfoCitiesGetRequest) SetIocType(iocType int64) error {
    r.iocType = iocType
    r.Set("ioc_type", iocType)
    return nil
}

// IocType Getter
func (r TaobaoAlitripTravelBaseinfoCitiesGetRequest) GetIocType() int64 {
    return r.iocType
}
// CatType Setter
// 1-境内跟团游 2-境内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 9-境内邮轮
func (r *TaobaoAlitripTravelBaseinfoCitiesGetRequest) SetCatType(catType int64) error {
    r.catType = catType
    r.Set("cat_type", catType)
    return nil
}

// CatType Getter
func (r TaobaoAlitripTravelBaseinfoCitiesGetRequest) GetCatType() int64 {
    return r.catType
}
