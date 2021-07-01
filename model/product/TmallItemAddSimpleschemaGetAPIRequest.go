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
type TmallItemAddSimpleschemaGetAPIRequest struct {
    model.Params
}

// 初始化TmallItemAddSimpleschemaGetAPIRequest对象
func NewTmallItemAddSimpleschemaGetRequest() *TmallItemAddSimpleschemaGetAPIRequest{
    return &TmallItemAddSimpleschemaGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemAddSimpleschemaGetAPIRequest) GetApiMethodName() string {
    return "tmall.item.add.simpleschema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemAddSimpleschemaGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
