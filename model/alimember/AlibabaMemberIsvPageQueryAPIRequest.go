package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberisvpagequeryAPIRequest isv离线会员数据分页查询 API请求
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
type AlibabamemberisvpagequeryAPIRequest struct {
	model.Params
	// 请求
	_pageQueryIsvCustomerRequest *PageQueryIsvCustomerRequest
}

// NewAlibabamemberisvpagequeryRequest 初始化AlibabamemberisvpagequeryAPIRequest对象
func NewAlibabamemberisvpagequeryRequest() *AlibabamemberisvpagequeryAPIRequest {
	return &AlibabamemberisvpagequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberisvpagequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.member.isv.page.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberisvpagequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberisvpagequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageQueryIsvCustomerRequest is PageQueryIsvCustomerRequest Setter
// 请求
func (r *AlibabamemberisvpagequeryAPIRequest) SetPageQueryIsvCustomerRequest(_pageQueryIsvCustomerRequest *PageQueryIsvCustomerRequest) error {
	r._pageQueryIsvCustomerRequest = _pageQueryIsvCustomerRequest
	r.Set("page_query_isv_customer_request", _pageQueryIsvCustomerRequest)
	return nil
}

// GetPageQueryIsvCustomerRequest PageQueryIsvCustomerRequest Getter
func (r AlibabamemberisvpagequeryAPIRequest) GetPageQueryIsvCustomerRequest() *PageQueryIsvCustomerRequest {
	return r._pageQueryIsvCustomerRequest
}
