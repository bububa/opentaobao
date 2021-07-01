package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗查询 API请求
alibaba.scbp.showcase.list

橱窗查询
*/
type AlibabaScbpShowcaseListAPIRequest struct {
    model.Params
    // 每页展示个数
    _perPageSize   int64
    // 页码
    _toPage   int64
}

// 初始化AlibabaScbpShowcaseListAPIRequest对象
func NewAlibabaScbpShowcaseListRequest() *AlibabaScbpShowcaseListAPIRequest{
    return &AlibabaScbpShowcaseListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseListAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PerPageSize Setter
// 每页展示个数
func (r *AlibabaScbpShowcaseListAPIRequest) SetPerPageSize(_perPageSize int64) error {
    r._perPageSize = _perPageSize
    r.Set("per_page_size", _perPageSize)
    return nil
}

// PerPageSize Getter
func (r AlibabaScbpShowcaseListAPIRequest) GetPerPageSize() int64 {
    return r._perPageSize
}
// ToPage Setter
// 页码
func (r *AlibabaScbpShowcaseListAPIRequest) SetToPage(_toPage int64) error {
    r._toPage = _toPage
    r.Set("to_page", _toPage)
    return nil
}

// ToPage Getter
func (r AlibabaScbpShowcaseListAPIRequest) GetToPage() int64 {
    return r._toPage
}
