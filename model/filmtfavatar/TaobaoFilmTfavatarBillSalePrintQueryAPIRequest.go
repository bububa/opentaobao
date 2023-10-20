package filmtfavatar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmtfavatarbillsaleprintqueryAPIRequest 获取影院卖品账单-核销账单 API请求
// taobao.film.tfavatar.bill.sale.print.query
//
// 获取影院卖品账单-核销账单
// 返回值data属于加密字段, 并非大字段.
type TaobaofilmtfavatarbillsaleprintqueryAPIRequest struct {
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

// NewTaobaofilmtfavatarbillsaleprintqueryRequest 初始化TaobaofilmtfavatarbillsaleprintqueryAPIRequest对象
func NewTaobaofilmtfavatarbillsaleprintqueryRequest() *TaobaofilmtfavatarbillsaleprintqueryAPIRequest {
	return &TaobaofilmtfavatarbillsaleprintqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfavatar.bill.sale.print.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIncludedOrderStatus is IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetIncludedOrderStatus(_includedOrderStatus []string) error {
	r._includedOrderStatus = _includedOrderStatus
	r.Set("included_order_status", _includedOrderStatus)
	return nil
}

// GetIncludedOrderStatus IncludedOrderStatus Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetIncludedOrderStatus() []string {
	return r._includedOrderStatus
}

// SetOpenAppKey is OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBeginTime is BeginTime Setter
// 开始时间
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCinemaId is CinemaId Setter
// 影院ID
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetCinemaId(_cinemaId int64) error {
	r._cinemaId = _cinemaId
	r.Set("cinema_id", _cinemaId)
	return nil
}

// GetCinemaId CinemaId Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetCinemaId() int64 {
	return r._cinemaId
}

// SetOffset is Offset Setter
// offset 下标, 从0开始
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaofilmtfavatarbillsaleprintqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofilmtfavatarbillsaleprintqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
