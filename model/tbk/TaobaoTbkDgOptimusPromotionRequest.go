package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-权益物料精选 APIRequest
taobao.tbk.dg.optimus.promotion

推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
*/
type TaobaoTbkDgOptimusPromotionRequest struct {
    model.Params

    // 页大小，一次请求请限制在10以内
    pageSize   int64 

    // 第几页，默认：1
    pageNum   int64 

    // mm_xxx_xxx_xxx的第3段数字
    adzoneId   int64 

    // 官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！
    promotionId   int64 

}

func NewTaobaoTbkDgOptimusPromotionRequest() *TaobaoTbkDgOptimusPromotionRequest{
    return &TaobaoTbkDgOptimusPromotionRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkDgOptimusPromotionRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.optimus.promotion"
}

func (r TaobaoTbkDgOptimusPromotionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkDgOptimusPromotionRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTbkDgOptimusPromotionRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTbkDgOptimusPromotionRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoTbkDgOptimusPromotionRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *TaobaoTbkDgOptimusPromotionRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

func (r TaobaoTbkDgOptimusPromotionRequest) GetAdzoneId() int64 {
    return r.adzoneId
}

func (r *TaobaoTbkDgOptimusPromotionRequest) SetPromotionId(promotionId int64) error {
    r.promotionId = promotionId
    r.Set("promotion_id", promotionId)
    return nil
}

func (r TaobaoTbkDgOptimusPromotionRequest) GetPromotionId() int64 {
    return r.promotionId
}

