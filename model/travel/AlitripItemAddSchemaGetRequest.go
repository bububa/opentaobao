package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布模板 APIRequest
alitrip.item.add.schema.get

发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemAddSchemaGetRequest struct {
    model.Params

    // 类目id
    catId   int64 

}

func NewAlitripItemAddSchemaGetRequest() *AlitripItemAddSchemaGetRequest{
    return &AlitripItemAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripItemAddSchemaGetRequest) GetApiMethodName() string {
    return "alitrip.item.add.schema.get"
}

func (r AlitripItemAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripItemAddSchemaGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlitripItemAddSchemaGetRequest) GetCatId() int64 {
    return r.catId
}

