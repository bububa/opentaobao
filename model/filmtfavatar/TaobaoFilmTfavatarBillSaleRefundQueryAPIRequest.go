package filmtfavatar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest
获取影院卖品账单--退款账单 API请求
taobao.film.tfavatar.bill.sale.refund.query

获取影院卖品账单--退款账单 */
type TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest struct {
	model.Params
	// 自运营开放平台APPKEY
	_openAppKey string
	// 影院ID
	_cinemaId int64
	// 开始时间
	_beginTime string
	// 结束时间
	_endTime string
	// 包含的订单状态, 默认不填
	_includedOrderStatus []string
	// offset 下标, 从0开始
	_offset int64
	// 页大小
	_pageSize int64
}

// NewTaobaoFilmTfavatarBillSaleRefundQueryRequest 初始化TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest对象
func NewTaobaoFilmTfavatarBillSaleRefundQueryRequest() *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest {
	return &TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfavatar.bill.sale.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// Get OpenAppKey Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// Set is CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetCinemaId(_cinemaId int64) error {
	r._cinemaId = _cinemaId
	r.Set("cinema_id", _cinemaId)
	return nil
}

// Get CinemaId Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetCinemaId() int64 {
	return r._cinemaId
}

// Set is BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// Get BeginTime Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// Set is EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetIncludedOrderStatus(_includedOrderStatus []string) error {
	r._includedOrderStatus = _includedOrderStatus
	r.Set("included_order_status", _includedOrderStatus)
	return nil
}

// Get IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetIncludedOrderStatus() []string {
	return r._includedOrderStatus
}

// Set is Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// Get Offset Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// Set is PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
