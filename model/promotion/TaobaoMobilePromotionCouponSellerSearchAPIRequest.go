package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionCouponSellerSearchAPIRequest 查询绑定卖家优惠券相关信息(手淘专用) API请求
// taobao.mobile.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
type TaobaoMobilePromotionCouponSellerSearchAPIRequest struct {
	model.Params
	// 请求id 排查线索 需保证单次调用唯一
	_traceId string
	// 券id集合
	_spreadIds string
	// 每页数据 最大20左右
	_pageSize int64
	// 当前第几页 从第一页开始
	_currentPage int64
}

// NewTaobaoMobilePromotionCouponSellerSearchRequest 初始化TaobaoMobilePromotionCouponSellerSearchAPIRequest对象
func NewTaobaoMobilePromotionCouponSellerSearchRequest() *TaobaoMobilePromotionCouponSellerSearchAPIRequest {
	return &TaobaoMobilePromotionCouponSellerSearchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMobilePromotionCouponSellerSearchAPIRequest) Reset() {
	r._traceId = ""
	r._spreadIds = ""
	r._pageSize = 0
	r._currentPage = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.coupon.seller.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 请求id 排查线索 需保证单次调用唯一
func (r *TaobaoMobilePromotionCouponSellerSearchAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetSpreadIds is SpreadIds Setter
// 券id集合
func (r *TaobaoMobilePromotionCouponSellerSearchAPIRequest) SetSpreadIds(_spreadIds string) error {
	r._spreadIds = _spreadIds
	r.Set("spread_ids", _spreadIds)
	return nil
}

// GetSpreadIds SpreadIds Getter
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetSpreadIds() string {
	return r._spreadIds
}

// SetPageSize is PageSize Setter
// 每页数据 最大20左右
func (r *TaobaoMobilePromotionCouponSellerSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前第几页 从第一页开始
func (r *TaobaoMobilePromotionCouponSellerSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoMobilePromotionCouponSellerSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

var poolTaobaoMobilePromotionCouponSellerSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMobilePromotionCouponSellerSearchRequest()
	},
}

// GetTaobaoMobilePromotionCouponSellerSearchRequest 从 sync.Pool 获取 TaobaoMobilePromotionCouponSellerSearchAPIRequest
func GetTaobaoMobilePromotionCouponSellerSearchAPIRequest() *TaobaoMobilePromotionCouponSellerSearchAPIRequest {
	return poolTaobaoMobilePromotionCouponSellerSearchAPIRequest.Get().(*TaobaoMobilePromotionCouponSellerSearchAPIRequest)
}

// ReleaseTaobaoMobilePromotionCouponSellerSearchAPIRequest 将 TaobaoMobilePromotionCouponSellerSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoMobilePromotionCouponSellerSearchAPIRequest(v *TaobaoMobilePromotionCouponSellerSearchAPIRequest) {
	v.Reset()
	poolTaobaoMobilePromotionCouponSellerSearchAPIRequest.Put(v)
}
