package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家在相关app领取的优惠券信息 APIRequest
taobao.promotion.coupon.buyer.search

查询买家在相关app领取的优惠券信息
*/
type TaobaoPromotionCouponBuyerSearchRequest struct {
    model.Params

    // 卖家昵称
    sellerNick   string 

    // 券状态  "正常",1 "已删除",-1 "已使用",-2 "冻结",0
    status   int64 

    // 每页数据 建议20左右
    pageSize   int64 

    // 当前第几页  从第一页开始
    currentPage   int64 

    // 结束时间
    endTime   string 

}

func NewTaobaoPromotionCouponBuyerSearchRequest() *TaobaoPromotionCouponBuyerSearchRequest{
    return &TaobaoPromotionCouponBuyerSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.buyer.search"
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionCouponBuyerSearchRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetSellerNick() string {
    return r.sellerNick
}

func (r *TaobaoPromotionCouponBuyerSearchRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoPromotionCouponBuyerSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoPromotionCouponBuyerSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoPromotionCouponBuyerSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionCouponBuyerSearchRequest) GetEndTime() string {
    return r.endTime
}

