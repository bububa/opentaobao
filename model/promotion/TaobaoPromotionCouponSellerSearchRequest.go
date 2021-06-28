package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息 APIRequest
taobao.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
*/
type TaobaoPromotionCouponSellerSearchRequest struct {
    model.Params

    // 卖家昵称
    sellerNick   string 

    // 当前第几页  从第一页开始
    currentPage   int64 

    // 每页数据 最大20左右
    pageSize   int64 

    // 券id集合
    spreadIds   []string 

}

func NewTaobaoPromotionCouponSellerSearchRequest() *TaobaoPromotionCouponSellerSearchRequest{
    return &TaobaoPromotionCouponSellerSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionCouponSellerSearchRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.seller.search"
}

func (r TaobaoPromotionCouponSellerSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionCouponSellerSearchRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoPromotionCouponSellerSearchRequest) GetSellerNick() string {
    return r.sellerNick
}

func (r *TaobaoPromotionCouponSellerSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoPromotionCouponSellerSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoPromotionCouponSellerSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPromotionCouponSellerSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoPromotionCouponSellerSearchRequest) SetSpreadIds(spreadIds []string) error {
    r.spreadIds = spreadIds
    r.Set("spread_ids", spreadIds)
    return nil
}

func (r TaobaoPromotionCouponSellerSearchRequest) GetSpreadIds() []string {
    return r.spreadIds
}

