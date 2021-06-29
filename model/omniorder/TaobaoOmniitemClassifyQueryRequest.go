package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分类信息 APIRequest
taobao.omniitem.classify.query

通过查询关键字，分页查询分类信息
*/
type TaobaoOmniitemClassifyQueryRequest struct {
    model.Params

    // 查询关键词
    keyword   string 

    // 页码
    pageNum   int64 

    // 每页大小
    pageSize   int64 

}

func NewTaobaoOmniitemClassifyQueryRequest() *TaobaoOmniitemClassifyQueryRequest{
    return &TaobaoOmniitemClassifyQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemClassifyQueryRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.query"
}

func (r TaobaoOmniitemClassifyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemClassifyQueryRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r TaobaoOmniitemClassifyQueryRequest) GetKeyword() string {
    return r.keyword
}

func (r *TaobaoOmniitemClassifyQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoOmniitemClassifyQueryRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *TaobaoOmniitemClassifyQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOmniitemClassifyQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

