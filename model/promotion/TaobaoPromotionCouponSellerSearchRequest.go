package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息 API请求
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

// 初始化TaobaoPromotionCouponSellerSearchRequest对象
func NewTaobaoPromotionCouponSellerSearchRequest() *TaobaoPromotionCouponSellerSearchRequest{
    return &TaobaoPromotionCouponSellerSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponSellerSearchRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.seller.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponSellerSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerNick Setter
// 卖家昵称
func (r *TaobaoPromotionCouponSellerSearchRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoPromotionCouponSellerSearchRequest) GetSellerNick() string {
    return r.sellerNick
}
// CurrentPage Setter
// 当前第几页  从第一页开始
func (r *TaobaoPromotionCouponSellerSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoPromotionCouponSellerSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 每页数据 最大20左右
func (r *TaobaoPromotionCouponSellerSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionCouponSellerSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
// SpreadIds Setter
// 券id集合
func (r *TaobaoPromotionCouponSellerSearchRequest) SetSpreadIds(spreadIds []string) error {
    r.spreadIds = spreadIds
    r.Set("spread_ids", spreadIds)
    return nil
}

// SpreadIds Getter
func (r TaobaoPromotionCouponSellerSearchRequest) GetSpreadIds() []string {
    return r.spreadIds
}
