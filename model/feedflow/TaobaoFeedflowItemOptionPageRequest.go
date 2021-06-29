package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询定向标签列表 APIRequest
taobao.feedflow.item.option.page

分页查询定向标签列表
*/
type TaobaoFeedflowItemOptionPageRequest struct {
    model.Params

    // 标签查询条件
    labelQuery   *LabelQueryDto 

}

func NewTaobaoFeedflowItemOptionPageRequest() *TaobaoFeedflowItemOptionPageRequest{
    return &TaobaoFeedflowItemOptionPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemOptionPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.option.page"
}

func (r TaobaoFeedflowItemOptionPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemOptionPageRequest) SetLabelQuery(labelQuery *LabelQueryDto) error {
    r.labelQuery = labelQuery
    r.Set("label_query", labelQuery)
    return nil
}

func (r TaobaoFeedflowItemOptionPageRequest) GetLabelQuery() *LabelQueryDto {
    return r.labelQuery
}

