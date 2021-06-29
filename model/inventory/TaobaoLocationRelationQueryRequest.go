package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地点关联关系查询 APIRequest
taobao.location.relation.query

地点关联关系查询 
门店和仓库关联关系查询
*/
type TaobaoLocationRelationQueryRequest struct {
    model.Params

    // 关系查询
    locationRelation   *LocationRelationDto 

}

func NewTaobaoLocationRelationQueryRequest() *TaobaoLocationRelationQueryRequest{
    return &TaobaoLocationRelationQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLocationRelationQueryRequest) GetApiMethodName() string {
    return "taobao.location.relation.query"
}

func (r TaobaoLocationRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLocationRelationQueryRequest) SetLocationRelation(locationRelation *LocationRelationDto) error {
    r.locationRelation = locationRelation
    r.Set("location_relation", locationRelation)
    return nil
}

func (r TaobaoLocationRelationQueryRequest) GetLocationRelation() *LocationRelationDto {
    return r.locationRelation
}

