package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponsGetAPIRequest 查询卖家优惠券 API请求
// taobao.promotion.coupons.get
//
// 查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量
type TaobaoPromotionCouponsGetAPIRequest struct {
	model.Params
	// 优惠券的截止日期
	_endTime string
	// 优惠券的id，唯一标识这个优惠券
	_couponId int64
	// 优惠券的面额，必须是3，5，10，20，50,100
	_denominations int64
	// 查询的页号，结果集是分页返回的，每页20条
	_pageNo int64
	// 每页条数
	_pageSize int64
}

// NewTaobaoPromotionCouponsGetRequest 初始化TaobaoPromotionCouponsGetAPIRequest对象
func NewTaobaoPromotionCouponsGetRequest() *TaobaoPromotionCouponsGetAPIRequest {
	return &TaobaoPromotionCouponsGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionCouponsGetAPIRequest) Reset() {
	r._endTime = ""
	r._couponId = 0
	r._denominations = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponsGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupons.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionCouponsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndTime is EndTime Setter
// 优惠券的截止日期
func (r *TaobaoPromotionCouponsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoPromotionCouponsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCouponId is CouponId Setter
// 优惠券的id，唯一标识这个优惠券
func (r *TaobaoPromotionCouponsGetAPIRequest) SetCouponId(_couponId int64) error {
	r._couponId = _couponId
	r.Set("coupon_id", _couponId)
	return nil
}

// GetCouponId CouponId Getter
func (r TaobaoPromotionCouponsGetAPIRequest) GetCouponId() int64 {
	return r._couponId
}

// SetDenominations is Denominations Setter
// 优惠券的面额，必须是3，5，10，20，50,100
func (r *TaobaoPromotionCouponsGetAPIRequest) SetDenominations(_denominations int64) error {
	r._denominations = _denominations
	r.Set("denominations", _denominations)
	return nil
}

// GetDenominations Denominations Getter
func (r TaobaoPromotionCouponsGetAPIRequest) GetDenominations() int64 {
	return r._denominations
}

// SetPageNo is PageNo Setter
// 查询的页号，结果集是分页返回的，每页20条
func (r *TaobaoPromotionCouponsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoPromotionCouponsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaoPromotionCouponsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPromotionCouponsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoPromotionCouponsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionCouponsGetRequest()
	},
}

// GetTaobaoPromotionCouponsGetRequest 从 sync.Pool 获取 TaobaoPromotionCouponsGetAPIRequest
func GetTaobaoPromotionCouponsGetAPIRequest() *TaobaoPromotionCouponsGetAPIRequest {
	return poolTaobaoPromotionCouponsGetAPIRequest.Get().(*TaobaoPromotionCouponsGetAPIRequest)
}

// ReleaseTaobaoPromotionCouponsGetAPIRequest 将 TaobaoPromotionCouponsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionCouponsGetAPIRequest(v *TaobaoPromotionCouponsGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionCouponsGetAPIRequest.Put(v)
}
