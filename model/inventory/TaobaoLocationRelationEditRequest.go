package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地点关联关系增量编辑 APIRequest
taobao.location.relation.edit

地点关联关系增量编辑
*/
type TaobaoLocationRelationEditRequest struct {
    model.Params

    // 关系对象列表
    locationRelationList   []LocationRelationDto 

}

func NewTaobaoLocationRelationEditRequest() *TaobaoLocationRelationEditRequest{
    return &TaobaoLocationRelationEditRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLocationRelationEditRequest) GetApiMethodName() string {
    return "taobao.location.relation.edit"
}

func (r TaobaoLocationRelationEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLocationRelationEditRequest) SetLocationRelationList(locationRelationList []LocationRelationDto) error {
    r.locationRelationList = locationRelationList
    r.Set("location_relation_list", locationRelationList)
    return nil
}

func (r TaobaoLocationRelationEditRequest) GetLocationRelationList() []LocationRelationDto {
    return r.locationRelationList
}

