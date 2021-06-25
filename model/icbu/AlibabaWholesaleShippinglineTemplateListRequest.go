package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运费模板 APIRequest
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

func NewAlibabaWholesaleShippinglineTemplateListRequest() *AlibabaWholesaleShippinglineTemplateListRequest{
    return &AlibabaWholesaleShippinglineTemplateListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWholesaleShippinglineTemplateListRequest) GetApiMethodName() string {
    return "alibaba.wholesale.shippingline.template.list"
}

func (r AlibabaWholesaleShippinglineTemplateListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWholesaleShippinglineTemplateListRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r AlibabaWholesaleShippinglineTemplateListRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *AlibabaWholesaleShippinglineTemplateListRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r AlibabaWholesaleShippinglineTemplateListRequest) GetCount() int64 {
    return r.count
}

