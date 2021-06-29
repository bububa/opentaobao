package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单元列表 APIRequest
taobao.feedflow.item.adgroup.page

通过计划id查询单元信息
*/
type TaobaoFeedflowItemAdgroupPageRequest struct {
    model.Params

    // 系统自动生成
    adgroupQuery   *AdgroupQueryDto 

}

func NewTaobaoFeedflowItemAdgroupPageRequest() *TaobaoFeedflowItemAdgroupPageRequest{
    return &TaobaoFeedflowItemAdgroupPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.page"
}

func (r TaobaoFeedflowItemAdgroupPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupPageRequest) SetAdgroupQuery(adgroupQuery *AdgroupQueryDto) error {
    r.adgroupQuery = adgroupQuery
    r.Set("adgroup_query", adgroupQuery)
    return nil
}

func (r TaobaoFeedflowItemAdgroupPageRequest) GetAdgroupQuery() *AdgroupQueryDto {
    return r.adgroupQuery
}

