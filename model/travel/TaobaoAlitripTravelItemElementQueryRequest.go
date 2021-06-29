package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素查询接口 API请求
taobao.alitrip.travel.item.element.query

提供资源元素查询接口，支持商家查询已经发布过的资源元素
*/
type TaobaoAlitripTravelItemElementQueryRequest struct {
    model.Params
    // 需要查询的资源元素列表，最大列表长度为50
    _outerIds   []string
}

// 初始化TaobaoAlitripTravelItemElementQueryRequest对象
func NewTaobaoAlitripTravelItemElementQueryRequest() *TaobaoAlitripTravelItemElementQueryRequest{
    return &TaobaoAlitripTravelItemElementQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemElementQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.element.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemElementQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterIds Setter
// 需要查询的资源元素列表，最大列表长度为50
func (r *TaobaoAlitripTravelItemElementQueryRequest) SetOuterIds(_outerIds []string) error {
    r._outerIds = _outerIds
    r.Set("outer_ids", _outerIds)
    return nil
}

// OuterIds Getter
func (r TaobaoAlitripTravelItemElementQueryRequest) GetOuterIds() []string {
    return r._outerIds
}
