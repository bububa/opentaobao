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
type AlibabaWholesaleShippinglineTemplateListRequest struct {
    model.Params
    // 第几页从1开始
    pageNum   int64
    // 每页返回的数据个数
    count   int64
}

// 初始化AlibabaWholesaleShippinglineTemplateListRequest对象
func NewAlibabaWholesaleShippinglineTemplateListRequest() *AlibabaWholesaleShippinglineTemplateListRequest{
    return &AlibabaWholesaleShippinglineTemplateListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleShippinglineTemplateListRequest) GetApiMethodName() string {
    return "alibaba.wholesale.shippingline.template.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleShippinglineTemplateListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNum Setter
// 第几页从1开始
func (r *AlibabaWholesaleShippinglineTemplateListRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaWholesaleShippinglineTemplateListRequest) GetPageNum() int64 {
    return r.pageNum
}
// Count Setter
// 每页返回的数据个数
func (r *AlibabaWholesaleShippinglineTemplateListRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r AlibabaWholesaleShippinglineTemplateListRequest) GetCount() int64 {
    return r.count
}
