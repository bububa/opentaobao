package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-权益物料精选 API请求
taobao.tbk.dg.optimus.promotion

推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
*/
type TaobaoTbkDgOptimusPromotionRequest struct {
    model.Params
    // 页大小，一次请求请限制在10以内
    _pageSize   int64
    // 第几页，默认：1
    _pageNum   int64
    // mm_xxx_xxx_xxx的第3段数字
    _adzoneId   int64
    // 官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！
    _promotionId   int64
}

// 初始化TaobaoTbkDgOptimusPromotionRequest对象
func NewTaobaoTbkDgOptimusPromotionRequest() *TaobaoTbkDgOptimusPromotionRequest{
    return &TaobaoTbkDgOptimusPromotionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgOptimusPromotionRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.optimus.promotion"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgOptimusPromotionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页大小，一次请求请限制在10以内
func (r *TaobaoTbkDgOptimusPromotionRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkDgOptimusPromotionRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNum Setter
// 第几页，默认：1
func (r *TaobaoTbkDgOptimusPromotionRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoTbkDgOptimusPromotionRequest) GetPageNum() int64 {
    return r._pageNum
}
// AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkDgOptimusPromotionRequest) SetAdzoneId(_adzoneId int64) error {
    r._adzoneId = _adzoneId
    r.Set("adzone_id", _adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkDgOptimusPromotionRequest) GetAdzoneId() int64 {
    return r._adzoneId
}
// PromotionId Setter
// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！
func (r *TaobaoTbkDgOptimusPromotionRequest) SetPromotionId(_promotionId int64) error {
    r._promotionId = _promotionId
    r.Set("promotion_id", _promotionId)
    return nil
}

// PromotionId Getter
func (r TaobaoTbkDgOptimusPromotionRequest) GetPromotionId() int64 {
    return r._promotionId
}
