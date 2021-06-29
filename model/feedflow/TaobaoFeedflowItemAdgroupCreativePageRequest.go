package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元下查看创意 APIRequest
taobao.feedflow.item.adgroup.creative.page

信息流单元下查看创意
*/
type TaobaoFeedflowItemAdgroupCreativePageRequest struct {
    model.Params

    // 绑定查询条件
    creativeBindQuery   *CreativeBindQueryDto 

}

func NewTaobaoFeedflowItemAdgroupCreativePageRequest() *TaobaoFeedflowItemAdgroupCreativePageRequest{
    return &TaobaoFeedflowItemAdgroupCreativePageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupCreativePageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.creative.page"
}

func (r TaobaoFeedflowItemAdgroupCreativePageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupCreativePageRequest) SetCreativeBindQuery(creativeBindQuery *CreativeBindQueryDto) error {
    r.creativeBindQuery = creativeBindQuery
    r.Set("creative_bind_query", creativeBindQuery)
    return nil
}

func (r TaobaoFeedflowItemAdgroupCreativePageRequest) GetCreativeBindQuery() *CreativeBindQueryDto {
    return r.creativeBindQuery
}

