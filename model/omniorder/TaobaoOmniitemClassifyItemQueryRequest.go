package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分类查商品信息 APIRequest
taobao.omniitem.classify.item.query

商家根据分类查商品
*/
type TaobaoOmniitemClassifyItemQueryRequest struct {
    model.Params

    // 分类ID
    classifyId   int64 

    // 页码
    pageNum   int64 

    // 每页大小
    pageSize   int64 

}

func NewTaobaoOmniitemClassifyItemQueryRequest() *TaobaoOmniitemClassifyItemQueryRequest{
    return &TaobaoOmniitemClassifyItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemClassifyItemQueryRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.item.query"
}

func (r TaobaoOmniitemClassifyItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemClassifyItemQueryRequest) SetClassifyId(classifyId int64) error {
    r.classifyId = classifyId
    r.Set("classify_id", classifyId)
    return nil
}

func (r TaobaoOmniitemClassifyItemQueryRequest) GetClassifyId() int64 {
    return r.classifyId
}

func (r *TaobaoOmniitemClassifyItemQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoOmniitemClassifyItemQueryRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *TaobaoOmniitemClassifyItemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOmniitemClassifyItemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

