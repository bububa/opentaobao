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
type TaobaoPromotionCouponSellerSearchAPIRequest struct {
    model.Params
    // 卖家昵称
    _sellerNick   string
    // 当前第几页  从第一页开始
    _currentPage   int64
    // 每页数据 最大20左右
    _pageSize   int64
    // 券id集合
    _spreadIds   []string
}

// 初始化TaobaoPromotionCouponSellerSearchAPIRequest对象
func NewTaobaoPromotionCouponSellerSearchRequest() *TaobaoPromotionCouponSellerSearchAPIRequest{
    return &TaobaoPromotionCouponSellerSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponSellerSearchAPIRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.seller.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponSellerSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerNick Setter
// 卖家昵称
func (r *TaobaoPromotionCouponSellerSearchAPIRequest) SetSellerNick(_sellerNick string) error {
    r._sellerNick = _sellerNick
    r.Set("seller_nick", _sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoPromotionCouponSellerSearchAPIRequest) GetSellerNick() string {
    return r._sellerNick
}
// CurrentPage Setter
// 当前第几页  从第一页开始
func (r *TaobaoPromotionCouponSellerSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoPromotionCouponSellerSearchAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页数据 最大20左右
func (r *TaobaoPromotionCouponSellerSearchAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionCouponSellerSearchAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// SpreadIds Setter
// 券id集合
func (r *TaobaoPromotionCouponSellerSearchAPIRequest) SetSpreadIds(_spreadIds []string) error {
    r._spreadIds = _spreadIds
    r.Set("spread_ids", _spreadIds)
    return nil
}

// SpreadIds Getter
func (r TaobaoPromotionCouponSellerSearchAPIRequest) GetSpreadIds() []string {
    return r._spreadIds
}
