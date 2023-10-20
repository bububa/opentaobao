package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaSmsNumQueryAPIRequest 短信查询 API请求
// alibaba.aliqin.ta.sms.num.query
//
// 查询短信发送揭露
type AlibabaAliqinTaSmsNumQueryAPIRequest struct {
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

// NewAlibabaAliqinTaSmsNumQueryRequest 初始化AlibabaAliqinTaSmsNumQueryAPIRequest对象
func NewAlibabaAliqinTaSmsNumQueryRequest() *AlibabaAliqinTaSmsNumQueryAPIRequest {
	return &AlibabaAliqinTaSmsNumQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) Reset() {
	r._bizId = ""
	r._recNum = ""
	r._queryDate = ""
	r._currentPage = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.sms.num.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 短信发送流水
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetBizId() string {
	return r._bizId
}

// SetRecNum is RecNum Setter
// 短信接收号码
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// GetRecNum RecNum Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetRecNum() string {
	return r._recNum
}

// SetQueryDate is QueryDate Setter
// 短信发送日期，支持近30天记录查询，格式yyyyMMdd
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetQueryDate(_queryDate string) error {
	r._queryDate = _queryDate
	r.Set("query_date", _queryDate)
	return nil
}

// GetQueryDate QueryDate Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetQueryDate() string {
	return r._queryDate
}

// SetCurrentPage is CurrentPage Setter
// 分页参数,页码
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 分页参数，每页数量。最大值50
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAliqinTaSmsNumQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinTaSmsNumQueryRequest()
	},
}

// GetAlibabaAliqinTaSmsNumQueryRequest 从 sync.Pool 获取 AlibabaAliqinTaSmsNumQueryAPIRequest
func GetAlibabaAliqinTaSmsNumQueryAPIRequest() *AlibabaAliqinTaSmsNumQueryAPIRequest {
	return poolAlibabaAliqinTaSmsNumQueryAPIRequest.Get().(*AlibabaAliqinTaSmsNumQueryAPIRequest)
}

// ReleaseAlibabaAliqinTaSmsNumQueryAPIRequest 将 AlibabaAliqinTaSmsNumQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinTaSmsNumQueryAPIRequest(v *AlibabaAliqinTaSmsNumQueryAPIRequest) {
	v.Reset()
	poolAlibabaAliqinTaSmsNumQueryAPIRequest.Put(v)
}
