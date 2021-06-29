package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询可用广告位列表 APIRequest
taobao.feedflow.item.adzone.list

批量查询可用广告位列表
*/
type TaobaoFeedflowItemAdzoneListRequest struct {
    model.Params

    // 广告位查询条件
    adzoneQuery   *AdzoneQueryDto 

}

func NewTaobaoFeedflowItemAdzoneListRequest() *TaobaoFeedflowItemAdzoneListRequest{
    return &TaobaoFeedflowItemAdzoneListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdzoneListRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.list"
}

func (r TaobaoFeedflowItemAdzoneListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdzoneListRequest) SetAdzoneQuery(adzoneQuery *AdzoneQueryDto) error {
    r.adzoneQuery = adzoneQuery
    r.Set("adzone_query", adzoneQuery)
    return nil
}

func (r TaobaoFeedflowItemAdzoneListRequest) GetAdzoneQuery() *AdzoneQueryDto {
    return r.adzoneQuery
}

