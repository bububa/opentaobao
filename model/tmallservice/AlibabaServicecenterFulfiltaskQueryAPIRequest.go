package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterFulfiltaskQueryAPIRequest 核销单查询 API请求
// alibaba.servicecenter.fulfiltask.query
//
// 当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
type AlibabaServicecenterFulfiltaskQueryAPIRequest struct {
	model.Params
	// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
	_gmtCreateStart string
	// 每页条数，默认50，最大50
	_pageSize int64
	// 核销单外部单号
	_outerId string
	// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
	_gmtCreateEnd string
	// 查询第几页
	_currentPage int64
	// 核销单id列表
	_fulfilTaskIdList []int64
}

// NewAlibabaServicecenterFulfiltaskQueryRequest 初始化AlibabaServicecenterFulfiltaskQueryAPIRequest对象
func NewAlibabaServicecenterFulfiltaskQueryRequest() *AlibabaServicecenterFulfiltaskQueryAPIRequest {
	return &AlibabaServicecenterFulfiltaskQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.fulfiltask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGmtCreateStart is GmtCreateStart Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaServicecenterFulfiltaskQueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetPageSize is PageSize Setter
// 每页条数，默认50，最大50
func (r *AlibabaServicecenterFulfiltaskQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOuterId is OuterId Setter
// 核销单外部单号
func (r *AlibabaServicecenterFulfiltaskQueryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaServicecenterFulfiltaskQueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetCurrentPage is CurrentPage Setter
// 查询第几页
func (r *AlibabaServicecenterFulfiltaskQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetFulfilTaskIdList is FulfilTaskIdList Setter
// 核销单id列表
func (r *AlibabaServicecenterFulfiltaskQueryAPIRequest) SetFulfilTaskIdList(_fulfilTaskIdList []int64) error {
	r._fulfilTaskIdList = _fulfilTaskIdList
	r.Set("fulfil_task_id_list", _fulfilTaskIdList)
	return nil
}

// GetFulfilTaskIdList FulfilTaskIdList Getter
func (r AlibabaServicecenterFulfiltaskQueryAPIRequest) GetFulfilTaskIdList() []int64 {
	return r._fulfilTaskIdList
}
