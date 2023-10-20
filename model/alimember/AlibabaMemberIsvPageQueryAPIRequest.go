package alimember

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIsvPageQueryAPIRequest isv离线会员数据分页查询 API请求
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
type AlibabaMemberIsvPageQueryAPIRequest struct {
	model.Params
	// 请求
	_pageQueryIsvCustomerRequest *PageQueryIsvCustomerRequest
}

// NewAlibabaMemberIsvPageQueryRequest 初始化AlibabaMemberIsvPageQueryAPIRequest对象
func NewAlibabaMemberIsvPageQueryRequest() *AlibabaMemberIsvPageQueryAPIRequest {
	return &AlibabaMemberIsvPageQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberIsvPageQueryAPIRequest) Reset() {
	r._pageQueryIsvCustomerRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberIsvPageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.member.isv.page.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberIsvPageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberIsvPageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageQueryIsvCustomerRequest is PageQueryIsvCustomerRequest Setter
// 请求
func (r *AlibabaMemberIsvPageQueryAPIRequest) SetPageQueryIsvCustomerRequest(_pageQueryIsvCustomerRequest *PageQueryIsvCustomerRequest) error {
	r._pageQueryIsvCustomerRequest = _pageQueryIsvCustomerRequest
	r.Set("page_query_isv_customer_request", _pageQueryIsvCustomerRequest)
	return nil
}

// GetPageQueryIsvCustomerRequest PageQueryIsvCustomerRequest Getter
func (r AlibabaMemberIsvPageQueryAPIRequest) GetPageQueryIsvCustomerRequest() *PageQueryIsvCustomerRequest {
	return r._pageQueryIsvCustomerRequest
}

var poolAlibabaMemberIsvPageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberIsvPageQueryRequest()
	},
}

// GetAlibabaMemberIsvPageQueryRequest 从 sync.Pool 获取 AlibabaMemberIsvPageQueryAPIRequest
func GetAlibabaMemberIsvPageQueryAPIRequest() *AlibabaMemberIsvPageQueryAPIRequest {
	return poolAlibabaMemberIsvPageQueryAPIRequest.Get().(*AlibabaMemberIsvPageQueryAPIRequest)
}

// ReleaseAlibabaMemberIsvPageQueryAPIRequest 将 AlibabaMemberIsvPageQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberIsvPageQueryAPIRequest(v *AlibabaMemberIsvPageQueryAPIRequest) {
	v.Reset()
	poolAlibabaMemberIsvPageQueryAPIRequest.Put(v)
}
