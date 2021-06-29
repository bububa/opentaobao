package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板列表 API请求
tmall.item.sizemapping.templates.list

获取所有尺码表模板列表。
*/
type TmallItemSizemappingTemplatesListRequest struct {
    model.Params
}

// 初始化TmallItemSizemappingTemplatesListRequest对象
func NewTmallItemSizemappingTemplatesListRequest() *TmallItemSizemappingTemplatesListRequest{
    return &TmallItemSizemappingTemplatesListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplatesListRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.templates.list"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplatesListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
