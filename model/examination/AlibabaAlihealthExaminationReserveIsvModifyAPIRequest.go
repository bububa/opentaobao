package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreserveisvmodifyAPIRequest ISV调TOP主动发起改期信息 API请求
// alibaba.alihealth.examination.reserve.isv.modify
//
// 体检机构对接_ISV发起体检套餐改期
type AlibabaalihealthexaminationreserveisvmodifyAPIRequest struct {
	model.Params
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
	// 修改后预约预约日期，格式：yyyy-MM-dd
	_reserveDate string
	// 修改后预约时间段开始时间 HH:mm:ss
	_reserveTimeStart string
	// 修改后预约时间段结束时间 HH:mm:ss
	_reserveTimeEnd string
}

// NewAlibabaalihealthexaminationreserveisvmodifyRequest 初始化AlibabaalihealthexaminationreserveisvmodifyAPIRequest对象
func NewAlibabaalihealthexaminationreserveisvmodifyRequest() *AlibabaalihealthexaminationreserveisvmodifyAPIRequest {
	return &AlibabaalihealthexaminationreserveisvmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.isv.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveNumber is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaalihealthexaminationreserveisvmodifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetUniqReserveCode is UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaalihealthexaminationreserveisvmodifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetReserveDate is ReserveDate Setter
// 修改后预约预约日期，格式：yyyy-MM-dd
func (r *AlibabaalihealthexaminationreserveisvmodifyAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// GetReserveDate ReserveDate Getter
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// SetReserveTimeStart is ReserveTimeStart Setter
// 修改后预约时间段开始时间 HH:mm:ss
func (r *AlibabaalihealthexaminationreserveisvmodifyAPIRequest) SetReserveTimeStart(_reserveTimeStart string) error {
	r._reserveTimeStart = _reserveTimeStart
	r.Set("reserve_time_start", _reserveTimeStart)
	return nil
}

// GetReserveTimeStart ReserveTimeStart Getter
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetReserveTimeStart() string {
	return r._reserveTimeStart
}

// SetReserveTimeEnd is ReserveTimeEnd Setter
// 修改后预约时间段结束时间 HH:mm:ss
func (r *AlibabaalihealthexaminationreserveisvmodifyAPIRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
	r._reserveTimeEnd = _reserveTimeEnd
	r.Set("reserve_time_end", _reserveTimeEnd)
	return nil
}

// GetReserveTimeEnd ReserveTimeEnd Getter
func (r AlibabaalihealthexaminationreserveisvmodifyAPIRequest) GetReserveTimeEnd() string {
	return r._reserveTimeEnd
}
