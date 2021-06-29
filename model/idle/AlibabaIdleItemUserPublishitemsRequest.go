package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布的商品列表 APIRequest
alibaba.idle.item.user.publishitems

为服务商的卖家提供发布的闲鱼商品列表
*/
type AlibabaIdleItemUserPublishitemsRequest struct {
    model.Params

    // 查询参数
    paramItemPageQuery   *ItemPageQuery 

}

func NewAlibabaIdleItemUserPublishitemsRequest() *AlibabaIdleItemUserPublishitemsRequest{
    return &AlibabaIdleItemUserPublishitemsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleItemUserPublishitemsRequest) GetApiMethodName() string {
    return "alibaba.idle.item.user.publishitems"
}

func (r AlibabaIdleItemUserPublishitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleItemUserPublishitemsRequest) SetParamItemPageQuery(paramItemPageQuery *ItemPageQuery) error {
    r.paramItemPageQuery = paramItemPageQuery
    r.Set("param_item_page_query", paramItemPageQuery)
    return nil
}

func (r AlibabaIdleItemUserPublishitemsRequest) GetParamItemPageQuery() *ItemPageQuery {
    return r.paramItemPageQuery
}

