package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌信息查询 API请求
alibaba.wdk.item.brand.query

品牌信息查询
*/
type AlibabaWdkItemBrandQueryRequest struct {
    model.Params
    // 查询关键词，不填则查询全部
    _keyword   string
    // 起始位置
    _offset   int64
    // 一页大小
    _pageSize   int64
}

// 初始化AlibabaWdkItemBrandQueryRequest对象
func NewAlibabaWdkItemBrandQueryRequest() *AlibabaWdkItemBrandQueryRequest{
    return &AlibabaWdkItemBrandQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemBrandQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.brand.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemBrandQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 查询关键词，不填则查询全部
func (r *AlibabaWdkItemBrandQueryRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlibabaWdkItemBrandQueryRequest) GetKeyword() string {
    return r._keyword
}
// Offset Setter
// 起始位置
func (r *AlibabaWdkItemBrandQueryRequest) SetOffset(_offset int64) error {
    r._offset = _offset
    r.Set("offset", _offset)
    return nil
}

// Offset Getter
func (r AlibabaWdkItemBrandQueryRequest) GetOffset() int64 {
    return r._offset
}
// PageSize Setter
// 一页大小
func (r *AlibabaWdkItemBrandQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkItemBrandQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
