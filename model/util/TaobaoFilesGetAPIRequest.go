package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilesgetAPIRequest 业务文件获取 API请求
// taobao.files.get
//
// 获取业务方暂存给ISV的文件列表
type TaobaofilesgetAPIRequest struct {
	model.Params
	// 搜索开始时间
	_startDate string
	// 搜索结束时间
	_endDate string
	// 下载链接状态。1:未下载。2:已下载
	_status int64
}

// NewTaobaofilesgetRequest 初始化TaobaofilesgetAPIRequest对象
func NewTaobaofilesgetRequest() *TaobaofilesgetAPIRequest {
	return &TaobaofilesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilesgetAPIRequest) GetApiMethodName() string {
	return "taobao.files.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 搜索开始时间
func (r *TaobaofilesgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaofilesgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 搜索结束时间
func (r *TaobaofilesgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaofilesgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStatus is Status Setter
// 下载链接状态。1:未下载。2:已下载
func (r *TaobaofilesgetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaofilesgetAPIRequest) GetStatus() int64 {
	return r._status
}
