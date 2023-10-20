package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterfulfiltaskqueryAPIRequest 核销单查询 API请求
// alibaba.servicecenter.fulfiltask.query
//
// 当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
type AlibabaservicecenterfulfiltaskqueryAPIRequest struct {
	model.Params
	// 核销单id列表
	_fulfilTaskIdList []int64
	// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
	_gmtCreateStart string
	// 核销单外部单号
	_outerId string
	// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
	_gmtCreateEnd string
	// 每页条数，默认50，最大50
	_pageSize int64
	// 查询第几页
	_currentPage int64
}

// NewAlibabaservicecenterfulfiltaskqueryRequest 初始化AlibabaservicecenterfulfiltaskqueryAPIRequest对象
func NewAlibabaservicecenterfulfiltaskqueryRequest() *AlibabaservicecenterfulfiltaskqueryAPIRequest {
	return &AlibabaservicecenterfulfiltaskqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.fulfiltask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfilTaskIdList is FulfilTaskIdList Setter
// 核销单id列表
func (r *AlibabaservicecenterfulfiltaskqueryAPIRequest) SetFulfilTaskIdList(_fulfilTaskIdList []int64) error {
	r._fulfilTaskIdList = _fulfilTaskIdList
	r.Set("fulfil_task_id_list", _fulfilTaskIdList)
	return nil
}

// GetFulfilTaskIdList FulfilTaskIdList Getter
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetFulfilTaskIdList() []int64 {
	return r._fulfilTaskIdList
}

// SetGmtCreateStart is GmtCreateStart Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaservicecenterfulfiltaskqueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetOuterId is OuterId Setter
// 核销单外部单号
func (r *AlibabaservicecenterfulfiltaskqueryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaservicecenterfulfiltaskqueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetPageSize is PageSize Setter
// 每页条数，默认50，最大50
func (r *AlibabaservicecenterfulfiltaskqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 查询第几页
func (r *AlibabaservicecenterfulfiltaskqueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaservicecenterfulfiltaskqueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
