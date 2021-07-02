package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilesGetAPIRequest 业务文件获取 API请求
// taobao.files.get
//
// 获取业务方暂存给ISV的文件列表
type TaobaoFilesGetAPIRequest struct {
	model.Params
	// 搜索开始时间
	_startDate string
	// 搜索结束时间
	_endDate string
	// 下载链接状态。1:未下载。2:已下载
	_status int64
}

// NewTaobaoFilesGetRequest 初始化TaobaoFilesGetAPIRequest对象
func NewTaobaoFilesGetRequest() *TaobaoFilesGetAPIRequest {
	return &TaobaoFilesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilesGetAPIRequest) GetApiMethodName() string {
	return "taobao.files.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartDate Setter
// 搜索开始时间
func (r *TaobaoFilesGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoFilesGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 搜索结束时间
func (r *TaobaoFilesGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoFilesGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is Status Setter
// 下载链接状态。1:未下载。2:已下载
func (r *TaobaoFilesGetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoFilesGetAPIRequest) GetStatus() int64 {
	return r._status
}
