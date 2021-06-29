package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元下查看绑定资源位 APIRequest
taobao.feedflow.item.adgroup.adzone.page

信息流单元下查看绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzonePageRequest struct {
    model.Params

    // 查询条件
    adzoneBindQuery   *AdzoneBindQueryDto 

}

func NewTaobaoFeedflowItemAdgroupAdzonePageRequest() *TaobaoFeedflowItemAdgroupAdzonePageRequest{
    return &TaobaoFeedflowItemAdgroupAdzonePageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupAdzonePageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.page"
}

func (r TaobaoFeedflowItemAdgroupAdzonePageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupAdzonePageRequest) SetAdzoneBindQuery(adzoneBindQuery *AdzoneBindQueryDto) error {
    r.adzoneBindQuery = adzoneBindQuery
    r.Set("adzone_bind_query", adzoneBindQuery)
    return nil
}

func (r TaobaoFeedflowItemAdgroupAdzonePageRequest) GetAdzoneBindQuery() *AdzoneBindQueryDto {
    return r.adzoneBindQuery
}

