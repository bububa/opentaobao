package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗查询 APIRequest
alibaba.scbp.showcase.list

橱窗查询
*/
type AlibabaScbpShowcaseListRequest struct {
    model.Params

    // 每页展示个数
    perPageSize   int64 

    // 页码
    toPage   int64 

}

func NewAlibabaScbpShowcaseListRequest() *AlibabaScbpShowcaseListRequest{
    return &AlibabaScbpShowcaseListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpShowcaseListRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.list"
}

func (r AlibabaScbpShowcaseListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpShowcaseListRequest) SetPerPageSize(perPageSize int64) error {
    r.perPageSize = perPageSize
    r.Set("per_page_size", perPageSize)
    return nil
}

func (r AlibabaScbpShowcaseListRequest) GetPerPageSize() int64 {
    return r.perPageSize
}

func (r *AlibabaScbpShowcaseListRequest) SetToPage(toPage int64) error {
    r.toPage = toPage
    r.Set("to_page", toPage)
    return nil
}

func (r AlibabaScbpShowcaseListRequest) GetToPage() int64 {
    return r.toPage
}

