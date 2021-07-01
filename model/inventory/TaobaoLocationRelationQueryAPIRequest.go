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
type TaobaoLocationRelationQueryAPIRequest struct {
    model.Params
    // 关系查询
    _locationRelation   *LocationRelationDto
}

// 初始化TaobaoLocationRelationQueryAPIRequest对象
func NewTaobaoLocationRelationQueryRequest() *TaobaoLocationRelationQueryAPIRequest{
    return &TaobaoLocationRelationQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLocationRelationQueryAPIRequest) GetApiMethodName() string {
    return "taobao.location.relation.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLocationRelationQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LocationRelation Setter
// 关系查询
func (r *TaobaoLocationRelationQueryAPIRequest) SetLocationRelation(_locationRelation *LocationRelationDto) error {
    r._locationRelation = _locationRelation
    r.Set("location_relation", _locationRelation)
    return nil
}

// LocationRelation Getter
func (r TaobaoLocationRelationQueryAPIRequest) GetLocationRelation() *LocationRelationDto {
    return r._locationRelation
}
