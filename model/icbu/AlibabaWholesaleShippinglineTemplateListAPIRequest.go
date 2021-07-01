package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运费模板 API请求
alibaba.wholesale.shippingline.template.list

查询运费模板信息
*/
type AlibabaWholesaleShippinglineTemplateListAPIRequest struct {
    model.Params
    // 第几页从1开始
    _pageNum   int64
    // 每页返回的数据个数
    _count   int64
}

// 初始化AlibabaWholesaleShippinglineTemplateListAPIRequest对象
func NewAlibabaWholesaleShippinglineTemplateListRequest() *AlibabaWholesaleShippinglineTemplateListAPIRequest{
    return &AlibabaWholesaleShippinglineTemplateListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetApiMethodName() string {
    return "alibaba.wholesale.shippingline.template.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNum Setter
// 第几页从1开始
func (r *AlibabaWholesaleShippinglineTemplateListAPIRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetPageNum() int64 {
    return r._pageNum
}
// Count Setter
// 每页返回的数据个数
func (r *AlibabaWholesaleShippinglineTemplateListAPIRequest) SetCount(_count int64) error {
    r._count = _count
    r.Set("count", _count)
    return nil
}

// Count Getter
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetCount() int64 {
    return r._count
}
