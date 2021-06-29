package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地点关联关系查询 API请求
taobao.location.relation.query

地点关联关系查询 
门店和仓库关联关系查询
*/
type TaobaoLocationRelationQueryRequest struct {
    model.Params
    // 关系查询
    locationRelation   *LocationRelationDto
}

// 初始化TaobaoLocationRelationQueryRequest对象
func NewTaobaoLocationRelationQueryRequest() *TaobaoLocationRelationQueryRequest{
    return &TaobaoLocationRelationQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLocationRelationQueryRequest) GetApiMethodName() string {
    return "taobao.location.relation.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLocationRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LocationRelation Setter
// 关系查询
func (r *TaobaoLocationRelationQueryRequest) SetLocationRelation(locationRelation *LocationRelationDto) error {
    r.locationRelation = locationRelation
    r.Set("location_relation", locationRelation)
    return nil
}

// LocationRelation Getter
func (r TaobaoLocationRelationQueryRequest) GetLocationRelation() *LocationRelationDto {
    return r.locationRelation
}
