package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息(手淘专用) API请求
taobao.mobile.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
*/
type TaobaoMobilePromotionCouponSellerSearchRequest struct {
    model.Params
    // 请求id 排查线索 需保证单次调用唯一
    _traceId   string
    // 券id集合
    _spreadIds   string
    // 每页数据 最大20左右
    _pageSize   int64
    // 当前第几页 从第一页开始
    _currentPage   int64
}

// 初始化TaobaoMobilePromotionCouponSellerSearchRequest对象
func NewTaobaoMobilePromotionCouponSellerSearchRequest() *TaobaoMobilePromotionCouponSellerSearchRequest{
    return &TaobaoMobilePromotionCouponSellerSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.coupon.seller.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceId Setter
// 请求id 排查线索 需保证单次调用唯一
func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetTraceId() string {
    return r._traceId
}
// SpreadIds Setter
// 券id集合
func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetSpreadIds(_spreadIds string) error {
    r._spreadIds = _spreadIds
    r.Set("spread_ids", _spreadIds)
    return nil
}

// SpreadIds Getter
func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetSpreadIds() string {
    return r._spreadIds
}
// PageSize Setter
// 每页数据 最大20左右
func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 当前第几页 从第一页开始
func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
