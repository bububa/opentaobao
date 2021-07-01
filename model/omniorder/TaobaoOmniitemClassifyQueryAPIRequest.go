package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分类信息 API请求
taobao.omniitem.classify.query

通过查询关键字，分页查询分类信息
*/
type TaobaoOmniitemClassifyQueryAPIRequest struct {
    model.Params
    // 查询关键词
    _keyword   string
    // 页码
    _pageNum   int64
    // 每页大小
    _pageSize   int64
}

// 初始化TaobaoOmniitemClassifyQueryAPIRequest对象
func NewTaobaoOmniitemClassifyQueryRequest() *TaobaoOmniitemClassifyQueryAPIRequest{
    return &TaobaoOmniitemClassifyQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyQueryAPIRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 查询关键词
func (r *TaobaoOmniitemClassifyQueryAPIRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r TaobaoOmniitemClassifyQueryAPIRequest) GetKeyword() string {
    return r._keyword
}
// PageNum Setter
// 页码
func (r *TaobaoOmniitemClassifyQueryAPIRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoOmniitemClassifyQueryAPIRequest) GetPageNum() int64 {
    return r._pageNum
}
// PageSize Setter
// 每页大小
func (r *TaobaoOmniitemClassifyQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOmniitemClassifyQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
