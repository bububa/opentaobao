package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
vip商家发布商品的获取规则接口 API请求
tmall.item.vip.add.schema.get

获取vip商家发布商品的规则
*/
type TmallItemVipAddSchemaGetRequest struct {
    model.Params
}

// 初始化TmallItemVipAddSchemaGetRequest对象
func NewTmallItemVipAddSchemaGetRequest() *TmallItemVipAddSchemaGetRequest{
    return &TmallItemVipAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemVipAddSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.vip.add.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemVipAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
