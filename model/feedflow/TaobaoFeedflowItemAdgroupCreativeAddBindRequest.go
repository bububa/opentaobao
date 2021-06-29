package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流新增并且绑定创意 APIRequest
taobao.feedflow.item.adgroup.creative.add.bind

信息流新增并且绑定创意
*/
type TaobaoFeedflowItemAdgroupCreativeAddBindRequest struct {
    model.Params

    // 新增绑定的创意，一次最多2个
    creativeBindList   []CreativeBindDto 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemAdgroupCreativeAddBindRequest() *TaobaoFeedflowItemAdgroupCreativeAddBindRequest{
    return &TaobaoFeedflowItemAdgroupCreativeAddBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.creative.add.bind"
}

func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupCreativeAddBindRequest) SetCreativeBindList(creativeBindList []CreativeBindDto) error {
    r.creativeBindList = creativeBindList
    r.Set("creative_bind_list", creativeBindList)
    return nil
}

func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetCreativeBindList() []CreativeBindDto {
    return r.creativeBindList
}

func (r *TaobaoFeedflowItemAdgroupCreativeAddBindRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

