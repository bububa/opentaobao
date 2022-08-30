package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest 根据商品id查询核销卷信息 API请求
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
type AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest struct {
	model.Params
	// 入参对象
	_queryDto *ProjectVerifyTicketQueryDto
}

// NewAlibabaAlihouseNewhomeProjectTicketQueryRequest 初始化AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectTicketQueryRequest() *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest {
	return &AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.ticket.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryDto is QueryDto Setter
// 入参对象
func (r *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) SetQueryDto(_queryDto *ProjectVerifyTicketQueryDto) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetQueryDto() *ProjectVerifyTicketQueryDto {
	return r._queryDto
}
