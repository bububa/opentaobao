package filmtfavatar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest 获取影院卖品账单--退款账单 API请求
// taobao.film.tfavatar.bill.sale.refund.query
//
// 获取影院卖品账单--退款账单
type TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest struct {
	model.Params
	// 包含的订单状态, 默认不填
	_includedOrderStatus []string
	// 自运营开放平台APPKEY
	_openAppKey string
	// 开始时间
	_beginTime string
	// 结束时间
	_endTime string
	// 影院ID
	_cinemaId int64
	// offset 下标, 从0开始
	_offset int64
	// 页大小
	_pageSize int64
}

// NewTaobaoFilmTfavatarBillSaleRefundQueryRequest 初始化TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest对象
func NewTaobaoFilmTfavatarBillSaleRefundQueryRequest() *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest {
	return &TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) Reset() {
	r._includedOrderStatus = r._includedOrderStatus[:0]
	r._openAppKey = ""
	r._beginTime = ""
	r._endTime = ""
	r._cinemaId = 0
	r._offset = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfavatar.bill.sale.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIncludedOrderStatus is IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetIncludedOrderStatus(_includedOrderStatus []string) error {
	r._includedOrderStatus = _includedOrderStatus
	r.Set("included_order_status", _includedOrderStatus)
	return nil
}

// GetIncludedOrderStatus IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetIncludedOrderStatus() []string {
	return r._includedOrderStatus
}

// SetOpenAppKey is OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBeginTime is BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCinemaId is CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetCinemaId(_cinemaId int64) error {
	r._cinemaId = _cinemaId
	r.Set("cinema_id", _cinemaId)
	return nil
}

// GetCinemaId CinemaId Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetCinemaId() int64 {
	return r._cinemaId
}

// SetOffset is Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoFilmTfavatarBillSaleRefundQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFilmTfavatarBillSaleRefundQueryRequest()
	},
}

// GetTaobaoFilmTfavatarBillSaleRefundQueryRequest 从 sync.Pool 获取 TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest
func GetTaobaoFilmTfavatarBillSaleRefundQueryAPIRequest() *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest {
	return poolTaobaoFilmTfavatarBillSaleRefundQueryAPIRequest.Get().(*TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest)
}

// ReleaseTaobaoFilmTfavatarBillSaleRefundQueryAPIRequest 将 TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoFilmTfavatarBillSaleRefundQueryAPIRequest(v *TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest) {
	v.Reset()
	poolTaobaoFilmTfavatarBillSaleRefundQueryAPIRequest.Put(v)
}
