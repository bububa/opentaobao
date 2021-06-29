package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 API请求
tmall.item.add.simpleschema.get

通过商家信息获取商品发布字段和规则。
*/
type TmallItemAddSimpleschemaGetRequest struct {
    model.Params
}

// 初始化TmallItemAddSimpleschemaGetRequest对象
func NewTmallItemAddSimpleschemaGetRequest() *TmallItemAddSimpleschemaGetRequest{
    return &TmallItemAddSimpleschemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemAddSimpleschemaGetRequest) GetApiMethodName() string {
    return "tmall.item.add.simpleschema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemAddSimpleschemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
