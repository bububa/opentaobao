package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponBuyerSearchAPIRequest 查询买家在相关app领取的优惠券信息 API请求
// taobao.promotion.coupon.buyer.search
//
// 查询买家在相关app领取的优惠券信息
type TaobaoPromotionCouponBuyerSearchAPIRequest struct {
	model.Params
	// 卖家昵称
	_sellerNick string
	// 结束时间
	_endTime string
	// 券状态  "正常",1 "已删除",-1 "已使用",-2 "冻结",0
	_status int64
	// 每页数据 建议20左右
	_pageSize int64
	// 当前第几页  从第一页开始
	_currentPage int64
}

// NewTaobaoPromotionCouponBuyerSearchRequest 初始化TaobaoPromotionCouponBuyerSearchAPIRequest对象
func NewTaobaoPromotionCouponBuyerSearchRequest() *TaobaoPromotionCouponBuyerSearchAPIRequest {
	return &TaobaoPromotionCouponBuyerSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.buyer.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSellerNick is SellerNick Setter
// 卖家昵称
func (r *TaobaoPromotionCouponBuyerSearchAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoPromotionCouponBuyerSearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// 券状态  "正常",1 "已删除",-1 "已使用",-2 "冻结",0
func (r *TaobaoPromotionCouponBuyerSearchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageSize is PageSize Setter
// 每页数据 建议20左右
func (r *TaobaoPromotionCouponBuyerSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前第几页  从第一页开始
func (r *TaobaoPromotionCouponBuyerSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoPromotionCouponBuyerSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
