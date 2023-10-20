package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqintasmsnumqueryAPIRequest 短信查询 API请求
// alibaba.aliqin.ta.sms.num.query
//
// 查询短信发送揭露
type AlibabaaliqintasmsnumqueryAPIRequest struct {
	model.Params
	// 短信发送流水
	_bizId string
	// 短信接收号码
	_recNum string
	// 短信发送日期，支持近30天记录查询，格式yyyyMMdd
	_queryDate string
	// 分页参数,页码
	_currentPage int64
	// 分页参数，每页数量。最大值50
	_pageSize int64
}

// NewAlibabaaliqintasmsnumqueryRequest 初始化AlibabaaliqintasmsnumqueryAPIRequest对象
func NewAlibabaaliqintasmsnumqueryRequest() *AlibabaaliqintasmsnumqueryAPIRequest {
	return &AlibabaaliqintasmsnumqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.sms.num.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 短信发送流水
func (r *AlibabaaliqintasmsnumqueryAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetBizId() string {
	return r._bizId
}

// SetRecNum is RecNum Setter
// 短信接收号码
func (r *AlibabaaliqintasmsnumqueryAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// GetRecNum RecNum Getter
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetRecNum() string {
	return r._recNum
}

// SetQueryDate is QueryDate Setter
// 短信发送日期，支持近30天记录查询，格式yyyyMMdd
func (r *AlibabaaliqintasmsnumqueryAPIRequest) SetQueryDate(_queryDate string) error {
	r._queryDate = _queryDate
	r.Set("query_date", _queryDate)
	return nil
}

// GetQueryDate QueryDate Getter
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetQueryDate() string {
	return r._queryDate
}

// SetCurrentPage is CurrentPage Setter
// 分页参数,页码
func (r *AlibabaaliqintasmsnumqueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 分页参数，每页数量。最大值50
func (r *AlibabaaliqintasmsnumqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaaliqintasmsnumqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
