package filmtfavatar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSalePrintQueryAPIRequest 获取影院卖品账单-核销账单 API请求
// taobao.film.tfavatar.bill.sale.print.query
//
// 获取影院卖品账单-核销账单
// 返回值data属于加密字段, 并非大字段.
type TaobaoFilmTfavatarBillSalePrintQueryAPIRequest struct {
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

// NewTaobaoFilmTfavatarBillSalePrintQueryRequest 初始化TaobaoFilmTfavatarBillSalePrintQueryAPIRequest对象
func NewTaobaoFilmTfavatarBillSalePrintQueryRequest() *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest {
	return &TaobaoFilmTfavatarBillSalePrintQueryAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) Reset() {
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
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfavatar.bill.sale.print.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIncludedOrderStatus is IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetIncludedOrderStatus(_includedOrderStatus []string) error {
	r._includedOrderStatus = _includedOrderStatus
	r.Set("included_order_status", _includedOrderStatus)
	return nil
}

// GetIncludedOrderStatus IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetIncludedOrderStatus() []string {
	return r._includedOrderStatus
}

// SetOpenAppKey is OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBeginTime is BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCinemaId is CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetCinemaId(_cinemaId int64) error {
	r._cinemaId = _cinemaId
	r.Set("cinema_id", _cinemaId)
	return nil
}

// GetCinemaId CinemaId Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetCinemaId() int64 {
	return r._cinemaId
}

// SetOffset is Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoFilmTfavatarBillSalePrintQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFilmTfavatarBillSalePrintQueryRequest()
	},
}

// GetTaobaoFilmTfavatarBillSalePrintQueryRequest 从 sync.Pool 获取 TaobaoFilmTfavatarBillSalePrintQueryAPIRequest
func GetTaobaoFilmTfavatarBillSalePrintQueryAPIRequest() *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest {
	return poolTaobaoFilmTfavatarBillSalePrintQueryAPIRequest.Get().(*TaobaoFilmTfavatarBillSalePrintQueryAPIRequest)
}

// ReleaseTaobaoFilmTfavatarBillSalePrintQueryAPIRequest 将 TaobaoFilmTfavatarBillSalePrintQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoFilmTfavatarBillSalePrintQueryAPIRequest(v *TaobaoFilmTfavatarBillSalePrintQueryAPIRequest) {
	v.Reset()
	poolTaobaoFilmTfavatarBillSalePrintQueryAPIRequest.Put(v)
}
