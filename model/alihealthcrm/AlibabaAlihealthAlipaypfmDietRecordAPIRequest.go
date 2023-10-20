package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthalipaypfmdietrecordAPIRequest 用户每日摄入卡路里总量回传接口 API请求
// alibaba.alihealth.alipaypfm.diet.record
//
// 用户每日摄入卡路里总量回传接口
type AlibabaalihealthalipaypfmdietrecordAPIRequest struct {
	model.Params
	// 记录日期，format：yyyy-MM-dd
	_date string
	// 用户健康ID
	_userId int64
	// 累积摄入卡路里
	_energy int64
}

// NewAlibabaalihealthalipaypfmdietrecordRequest 初始化AlibabaalihealthalipaypfmdietrecordAPIRequest对象
func NewAlibabaalihealthalipaypfmdietrecordRequest() *AlibabaalihealthalipaypfmdietrecordAPIRequest {
	return &AlibabaalihealthalipaypfmdietrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthalipaypfmdietrecordAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alipaypfm.diet.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthalipaypfmdietrecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthalipaypfmdietrecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 记录日期，format：yyyy-MM-dd
func (r *AlibabaalihealthalipaypfmdietrecordAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r AlibabaalihealthalipaypfmdietrecordAPIRequest) GetDate() string {
	return r._date
}

// SetUserId is UserId Setter
// 用户健康ID
func (r *AlibabaalihealthalipaypfmdietrecordAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalihealthalipaypfmdietrecordAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetEnergy is Energy Setter
// 累积摄入卡路里
func (r *AlibabaalihealthalipaypfmdietrecordAPIRequest) SetEnergy(_energy int64) error {
	r._energy = _energy
	r.Set("energy", _energy)
	return nil
}

// GetEnergy Energy Getter
func (r AlibabaalihealthalipaypfmdietrecordAPIRequest) GetEnergy() int64 {
	return r._energy
}
