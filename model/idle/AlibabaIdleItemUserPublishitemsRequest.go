package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布的商品列表 API请求
alibaba.idle.item.user.publishitems

为服务商的卖家提供发布的闲鱼商品列表
*/
type AlibabaIdleItemUserPublishitemsRequest struct {
    model.Params
    // 查询参数
    paramItemPageQuery   *ItemPageQuery
}

// 初始化AlibabaIdleItemUserPublishitemsRequest对象
func NewAlibabaIdleItemUserPublishitemsRequest() *AlibabaIdleItemUserPublishitemsRequest{
    return &AlibabaIdleItemUserPublishitemsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemUserPublishitemsRequest) GetApiMethodName() string {
    return "alibaba.idle.item.user.publishitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemUserPublishitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamItemPageQuery Setter
// 查询参数
func (r *AlibabaIdleItemUserPublishitemsRequest) SetParamItemPageQuery(paramItemPageQuery *ItemPageQuery) error {
    r.paramItemPageQuery = paramItemPageQuery
    r.Set("param_item_page_query", paramItemPageQuery)
    return nil
}

// ParamItemPageQuery Getter
func (r AlibabaIdleItemUserPublishitemsRequest) GetParamItemPageQuery() *ItemPageQuery {
    return r.paramItemPageQuery
}
