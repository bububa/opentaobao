package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌信息查询 APIRequest
alibaba.wdk.item.brand.query

品牌信息查询
*/
type AlibabaWdkItemBrandQueryRequest struct {
    model.Params

    // 查询关键词，不填则查询全部
    keyword   string 

    // 起始位置
    offset   int64 

    // 一页大小
    pageSize   int64 

}

func NewAlibabaWdkItemBrandQueryRequest() *AlibabaWdkItemBrandQueryRequest{
    return &AlibabaWdkItemBrandQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemBrandQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.brand.query"
}

func (r AlibabaWdkItemBrandQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemBrandQueryRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlibabaWdkItemBrandQueryRequest) GetKeyword() string {
    return r.keyword
}

func (r *AlibabaWdkItemBrandQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r AlibabaWdkItemBrandQueryRequest) GetOffset() int64 {
    return r.offset
}

func (r *AlibabaWdkItemBrandQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaWdkItemBrandQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

