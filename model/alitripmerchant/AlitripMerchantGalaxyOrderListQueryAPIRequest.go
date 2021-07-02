package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderListQueryAPIRequest 星河-订单列表查询 API请求
// alitrip.merchant.galaxy.order.list.query
//
// 为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
type AlitripMerchantGalaxyOrderListQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户token
	_token string
	// 订单状态
	_orderStatus string
	// 入住时间
	_startTime string
	// 入住时间
	_endTime string
	// 页数
	_page int64
	// 每页行数
	_row int64
}

// NewAlitripMerchantGalaxyOrderListQueryRequest 初始化AlitripMerchantGalaxyOrderListQueryAPIRequest对象
func NewAlitripMerchantGalaxyOrderListQueryRequest() *AlitripMerchantGalaxyOrderListQueryAPIRequest {
	return &AlitripMerchantGalaxyOrderListQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderStatus is OrderStatus Setter
// 订单状态
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetOrderStatus(_orderStatus string) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetOrderStatus() string {
	return r._orderStatus
}

// SetStartTime is StartTime Setter
// 入住时间
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 入住时间
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPage is Page Setter
// 页数
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetRow is Row Setter
// 每页行数
func (r *AlitripMerchantGalaxyOrderListQueryAPIRequest) SetRow(_row int64) error {
	r._row = _row
	r.Set("row", _row)
	return nil
}

// GetRow Row Getter
func (r AlitripMerchantGalaxyOrderListQueryAPIRequest) GetRow() int64 {
	return r._row
}
