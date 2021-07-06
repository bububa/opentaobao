package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoutonActivityListAPIRequest 商家自运营活动列表 API请求
// alibaba.mouton.activity.list
//
// 商家查询自己配置的活动列表
type AlibabaMoutonActivityListAPIRequest struct {
	model.Params
	// 状态
	_statusList []string
	// 开始时间
	_startTimeEnd string
	// 来源
	_source string
	// 开始时间
	_startTimeBegin string
	// 结束时间
	_endTimeBegin string
	// 结束时间
	_endTimeEnd string
	// 每页记录数
	_pageSize int64
	// 来源记录id
	_sourceRecordId int64
	// 当前页
	_currentPage int64
}

// NewAlibabaMoutonActivityListRequest 初始化AlibabaMoutonActivityListAPIRequest对象
func NewAlibabaMoutonActivityListRequest() *AlibabaMoutonActivityListAPIRequest {
	return &AlibabaMoutonActivityListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoutonActivityListAPIRequest) GetApiMethodName() string {
	return "alibaba.mouton.activity.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoutonActivityListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStatusList is StatusList Setter
// 状态
func (r *AlibabaMoutonActivityListAPIRequest) SetStatusList(_statusList []string) error {
	r._statusList = _statusList
	r.Set("status_list", _statusList)
	return nil
}

// GetStatusList StatusList Getter
func (r AlibabaMoutonActivityListAPIRequest) GetStatusList() []string {
	return r._statusList
}

// SetStartTimeEnd is StartTimeEnd Setter
// 开始时间
func (r *AlibabaMoutonActivityListAPIRequest) SetStartTimeEnd(_startTimeEnd string) error {
	r._startTimeEnd = _startTimeEnd
	r.Set("start_time_end", _startTimeEnd)
	return nil
}

// GetStartTimeEnd StartTimeEnd Getter
func (r AlibabaMoutonActivityListAPIRequest) GetStartTimeEnd() string {
	return r._startTimeEnd
}

// SetSource is Source Setter
// 来源
func (r *AlibabaMoutonActivityListAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaMoutonActivityListAPIRequest) GetSource() string {
	return r._source
}

// SetStartTimeBegin is StartTimeBegin Setter
// 开始时间
func (r *AlibabaMoutonActivityListAPIRequest) SetStartTimeBegin(_startTimeBegin string) error {
	r._startTimeBegin = _startTimeBegin
	r.Set("start_time_begin", _startTimeBegin)
	return nil
}

// GetStartTimeBegin StartTimeBegin Getter
func (r AlibabaMoutonActivityListAPIRequest) GetStartTimeBegin() string {
	return r._startTimeBegin
}

// SetEndTimeBegin is EndTimeBegin Setter
// 结束时间
func (r *AlibabaMoutonActivityListAPIRequest) SetEndTimeBegin(_endTimeBegin string) error {
	r._endTimeBegin = _endTimeBegin
	r.Set("end_time_begin", _endTimeBegin)
	return nil
}

// GetEndTimeBegin EndTimeBegin Getter
func (r AlibabaMoutonActivityListAPIRequest) GetEndTimeBegin() string {
	return r._endTimeBegin
}

// SetEndTimeEnd is EndTimeEnd Setter
// 结束时间
func (r *AlibabaMoutonActivityListAPIRequest) SetEndTimeEnd(_endTimeEnd string) error {
	r._endTimeEnd = _endTimeEnd
	r.Set("end_time_end", _endTimeEnd)
	return nil
}

// GetEndTimeEnd EndTimeEnd Getter
func (r AlibabaMoutonActivityListAPIRequest) GetEndTimeEnd() string {
	return r._endTimeEnd
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *AlibabaMoutonActivityListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaMoutonActivityListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetSourceRecordId is SourceRecordId Setter
// 来源记录id
func (r *AlibabaMoutonActivityListAPIRequest) SetSourceRecordId(_sourceRecordId int64) error {
	r._sourceRecordId = _sourceRecordId
	r.Set("source_record_id", _sourceRecordId)
	return nil
}

// GetSourceRecordId SourceRecordId Getter
func (r AlibabaMoutonActivityListAPIRequest) GetSourceRecordId() int64 {
	return r._sourceRecordId
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlibabaMoutonActivityListAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaMoutonActivityListAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
