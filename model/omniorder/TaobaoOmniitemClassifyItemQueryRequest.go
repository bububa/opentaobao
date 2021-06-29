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
    _classifyId   int64
    // 页码
    _pageNum   int64
    // 每页大小
    _pageSize   int64
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
func (r *TaobaoOmniitemClassifyItemQueryRequest) SetClassifyId(_classifyId int64) error {
    r._classifyId = _classifyId
    r.Set("classify_id", _classifyId)
    return nil
}

// ClassifyId Getter
func (r TaobaoOmniitemClassifyItemQueryRequest) GetClassifyId() int64 {
    return r._classifyId
}
// PageNum Setter
// 页码
func (r *TaobaoOmniitemClassifyItemQueryRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoOmniitemClassifyItemQueryRequest) GetPageNum() int64 {
    return r._pageNum
}
// PageSize Setter
// 每页大小
func (r *TaobaoOmniitemClassifyItemQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOmniitemClassifyItemQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
