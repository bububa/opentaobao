package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户宝贝详情页模板名称 API请求
taobao.item.templates.get

查询当前登录用户的店铺的宝贝详情页的模板名称
*/
type TaobaoItemTemplatesGetRequest struct {
    model.Params
}

// 初始化TaobaoItemTemplatesGetRequest对象
func NewTaobaoItemTemplatesGetRequest() *TaobaoItemTemplatesGetRequest{
    return &TaobaoItemTemplatesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemTemplatesGetRequest) GetApiMethodName() string {
    return "taobao.item.templates.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemTemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
