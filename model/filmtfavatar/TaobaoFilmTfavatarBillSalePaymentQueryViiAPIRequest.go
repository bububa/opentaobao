package filmtfavatar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest 获取影院卖品账单--支付账单-V2版本(正逆分离) API请求
// taobao.film.tfavatar.bill.sale.payment.query.vii
//
// 获取影院卖品账单--支付账单-V2版本(正逆分离)
type TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest struct {
	model.Params
	// 自运营开放平台APPKEY
	_openAppKey string
	// 开始时间
	_beginTime string
	// 结束时间
	_endTime string
	// 包含的订单状态, 默认不填
	_includedOrderStatus string
	// 影院ID
	_cinemaId int64
	// offset 下标, 从0开始
	_offset int64
	// 页大小
	_pageSize int64
}

// NewTaobaoFilmTfavatarBillSalePaymentQueryViiRequest 初始化TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest对象
func NewTaobaoFilmTfavatarBillSalePaymentQueryViiRequest() *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest {
	return &TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) Reset() {
	r._openAppKey = ""
	r._beginTime = ""
	r._endTime = ""
	r._includedOrderStatus = ""
	r._cinemaId = 0
	r._offset = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfavatar.bill.sale.payment.query.vii"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAppKey is OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBeginTime is BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetIncludedOrderStatus is IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetIncludedOrderStatus(_includedOrderStatus string) error {
	r._includedOrderStatus = _includedOrderStatus
	r.Set("included_order_status", _includedOrderStatus)
	return nil
}

// GetIncludedOrderStatus IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetIncludedOrderStatus() string {
	return r._includedOrderStatus
}

// SetCinemaId is CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetCinemaId(_cinemaId int64) error {
	r._cinemaId = _cinemaId
	r.Set("cinema_id", _cinemaId)
	return nil
}

// GetCinemaId CinemaId Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetCinemaId() int64 {
	return r._cinemaId
}

// SetOffset is Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFilmTfavatarBillSalePaymentQueryViiRequest()
	},
}

// GetTaobaoFilmTfavatarBillSalePaymentQueryViiRequest 从 sync.Pool 获取 TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest
func GetTaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest() *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest {
	return poolTaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest.Get().(*TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest)
}

// ReleaseTaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest 将 TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest 放入 sync.Pool
func ReleaseTaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest(v *TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest) {
	v.Reset()
	poolTaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest.Put(v)
}
