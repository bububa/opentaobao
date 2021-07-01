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
type TaobaoItemTemplatesGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoItemTemplatesGetAPIRequest对象
func NewTaobaoItemTemplatesGetRequest() *TaobaoItemTemplatesGetAPIRequest{
    return &TaobaoItemTemplatesGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemTemplatesGetAPIRequest) GetApiMethodName() string {
    return "taobao.item.templates.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemTemplatesGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
