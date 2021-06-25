package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素查询接口 APIRequest
taobao.alitrip.travel.item.element.query

提供资源元素查询接口，支持商家查询已经发布过的资源元素
*/
type TaobaoAlitripTravelItemElementQueryRequest struct {
    model.Params

    // 需要查询的资源元素列表，最大列表长度为50
    outerIds   []String 

}

func NewTaobaoAlitripTravelItemElementQueryRequest() *TaobaoAlitripTravelItemElementQueryRequest{
    return &TaobaoAlitripTravelItemElementQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelItemElementQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.element.query"
}

func (r TaobaoAlitripTravelItemElementQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelItemElementQueryRequest) SetOuterIds(outerIds []String) error {
    r.outerIds = outerIds
    r.Set("outer_ids", outerIds)
    return nil
}

func (r TaobaoAlitripTravelItemElementQueryRequest) GetOuterIds() []String {
    return r.outerIds
}

