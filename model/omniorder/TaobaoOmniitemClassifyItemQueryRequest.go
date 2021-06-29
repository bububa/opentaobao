package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分类查商品信息 API请求
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

// 初始化TaobaoOmniitemClassifyItemQueryRequest对象
func NewTaobaoOmniitemClassifyItemQueryRequest() *TaobaoOmniitemClassifyItemQueryRequest{
    return &TaobaoOmniitemClassifyItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyItemQueryRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.item.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClassifyId Setter
// 分类ID
func (r *TaobaoOmniitemClassifyItemQueryRequest) SetClassifyId(classifyId int64) error {
    r.classifyId = classifyId
    r.Set("classify_id", classifyId)
    return nil
}

// ClassifyId Getter
func (r TaobaoOmniitemClassifyItemQueryRequest) GetClassifyId() int64 {
    return r.classifyId
}
// PageNum Setter
// 页码
func (r *TaobaoOmniitemClassifyItemQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoOmniitemClassifyItemQueryRequest) GetPageNum() int64 {
    return r.pageNum
}
// PageSize Setter
// 每页大小
func (r *TaobaoOmniitemClassifyItemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOmniitemClassifyItemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
