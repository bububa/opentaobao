package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流查看商品列表 API请求
taobao.feedflow.item.item.page

信息流查看商品列表
*/
type TaobaoFeedflowItemItemPageRequest struct {
    model.Params
    // 查询条件
    _itemQuery   *ItemQueryDTO
}

// 初始化TaobaoFeedflowItemItemPageRequest对象
func NewTaobaoFeedflowItemItemPageRequest() *TaobaoFeedflowItemItemPageRequest{
    return &TaobaoFeedflowItemItemPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemItemPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.item.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemItemPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemItemPageRequest) SetItemQuery(_itemQuery *ItemQueryDTO) error {
    r._itemQuery = _itemQuery
    r.Set("item_query", _itemQuery)
    return nil
}

// ItemQuery Getter
func (r TaobaoFeedflowItemItemPageRequest) GetItemQuery() *ItemQueryDTO {
    return r._itemQuery
}
