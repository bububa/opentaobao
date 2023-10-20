package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectticketqueryAPIRequest 根据商品id查询核销卷信息 API请求
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
type AlibabaalihousenewhomeprojectticketqueryAPIRequest struct {
	model.Params
	// 入参对象
	_queryDto *ProjectVerifyTicketQueryDto
}

// NewAlibabaalihousenewhomeprojectticketqueryRequest 初始化AlibabaalihousenewhomeprojectticketqueryAPIRequest对象
func NewAlibabaalihousenewhomeprojectticketqueryRequest() *AlibabaalihousenewhomeprojectticketqueryAPIRequest {
	return &AlibabaalihousenewhomeprojectticketqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectticketqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.ticket.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectticketqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectticketqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// 入参对象
func (r *AlibabaalihousenewhomeprojectticketqueryAPIRequest) SetQueryDto(_queryDto *ProjectVerifyTicketQueryDto) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaalihousenewhomeprojectticketqueryAPIRequest) GetQueryDto() *ProjectVerifyTicketQueryDto {
	return r._queryDto
}
