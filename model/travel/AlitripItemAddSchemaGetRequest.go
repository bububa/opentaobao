package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布模板 API请求
alitrip.item.add.schema.get

发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemAddSchemaGetRequest struct {
    model.Params
    // 类目id
    catId   int64
}

// 初始化AlitripItemAddSchemaGetRequest对象
func NewAlitripItemAddSchemaGetRequest() *AlitripItemAddSchemaGetRequest{
    return &AlitripItemAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripItemAddSchemaGetRequest) GetApiMethodName() string {
    return "alitrip.item.add.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripItemAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *AlitripItemAddSchemaGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlitripItemAddSchemaGetRequest) GetCatId() int64 {
    return r.catId
}
