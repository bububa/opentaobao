package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流查看商品列表 APIRequest
taobao.feedflow.item.item.page

信息流查看商品列表
*/
type TaobaoFeedflowItemItemPageRequest struct {
    model.Params

    // 查询条件
    itemQuery   *ItemQueryDto 

}

func NewTaobaoFeedflowItemItemPageRequest() *TaobaoFeedflowItemItemPageRequest{
    return &TaobaoFeedflowItemItemPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemItemPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.item.page"
}

func (r TaobaoFeedflowItemItemPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemItemPageRequest) SetItemQuery(itemQuery *ItemQueryDto) error {
    r.itemQuery = itemQuery
    r.Set("item_query", itemQuery)
    return nil
}

func (r TaobaoFeedflowItemItemPageRequest) GetItemQuery() *ItemQueryDto {
    return r.itemQuery
}

