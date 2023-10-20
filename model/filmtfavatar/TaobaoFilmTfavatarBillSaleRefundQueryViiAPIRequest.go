package filmtfavatar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest 获取影院卖品账单--退款账单-V2版本(正逆分离) API请求
// taobao.film.tfavatar.bill.sale.refund.query.vii
//
// 获取影院卖品账单--退款账单-V2版本(正逆分离)
type TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest struct {
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

// NewTaobaofilmtfavatarbillsalerefundqueryviiRequest 初始化TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest对象
func NewTaobaofilmtfavatarbillsalerefundqueryviiRequest() *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest {
	return &TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfavatar.bill.sale.refund.query.vii"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAppKey is OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBeginTime is BeginTime Setter
// 开始时间
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetIncludedOrderStatus is IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetIncludedOrderStatus(_includedOrderStatus string) error {
	r._includedOrderStatus = _includedOrderStatus
	r.Set("included_order_status", _includedOrderStatus)
	return nil
}

// GetIncludedOrderStatus IncludedOrderStatus Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetIncludedOrderStatus() string {
	return r._includedOrderStatus
}

// SetCinemaId is CinemaId Setter
// 影院ID
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetCinemaId(_cinemaId int64) error {
	r._cinemaId = _cinemaId
	r.Set("cinema_id", _cinemaId)
	return nil
}

// GetCinemaId CinemaId Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetCinemaId() int64 {
	return r._cinemaId
}

// SetOffset is Offset Setter
// offset 下标, 从0开始
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
