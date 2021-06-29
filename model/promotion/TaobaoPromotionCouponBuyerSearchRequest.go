package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家在相关app领取的优惠券信息 API请求
taobao.promotion.coupon.buyer.search

查询买家在相关app领取的优惠券信息
*/
type TaobaoPromotionCouponBuyerSearchRequest struct {
    model.Params
    // 卖家昵称
    _sellerNick   string
    // 券状态  "正常",1 "已删除",-1 "已使用",-2 "冻结",0
    _status   int64
    // 每页数据 建议20左右
    _pageSize   int64
    // 当前第几页  从第一页开始
    _currentPage   int64
    // 结束时间
    _endTime   string
}

// 初始化TaobaoPromotionCouponBuyerSearchRequest对象
func NewTaobaoPromotionCouponBuyerSearchRequest() *TaobaoPromotionCouponBuyerSearchRequest{
    return &TaobaoPromotionCouponBuyerSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponBuyerSearchRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.buyer.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponBuyerSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerNick Setter
// 卖家昵称
func (r *TaobaoPromotionCouponBuyerSearchRequest) SetSellerNick(_sellerNick string) error {
    r._sellerNick = _sellerNick
    r.Set("seller_nick", _sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoPromotionCouponBuyerSearchRequest) GetSellerNick() string {
    return r._sellerNick
}
// Status Setter
// 券状态  "正常",1 "已删除",-1 "已使用",-2 "冻结",0
func (r *TaobaoPromotionCouponBuyerSearchRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoPromotionCouponBuyerSearchRequest) GetStatus() int64 {
    return r._status
}
// PageSize Setter
// 每页数据 建议20左右
func (r *TaobaoPromotionCouponBuyerSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionCouponBuyerSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 当前第几页  从第一页开始
func (r *TaobaoPromotionCouponBuyerSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoPromotionCouponBuyerSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// EndTime Setter
// 结束时间
func (r *TaobaoPromotionCouponBuyerSearchRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionCouponBuyerSearchRequest) GetEndTime() string {
    return r._endTime
}
