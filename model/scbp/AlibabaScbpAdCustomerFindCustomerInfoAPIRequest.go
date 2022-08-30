package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCustomerFindCustomerInfoAPIRequest 查询客户信息 API请求
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
type AlibabaScbpAdCustomerFindCustomerInfoAPIRequest struct {
	model.Params
	// 用户信息
	_topContextDto *TopContextDto
}

// NewAlibabaScbpAdCustomerFindCustomerInfoRequest 初始化AlibabaScbpAdCustomerFindCustomerInfoAPIRequest对象
func NewAlibabaScbpAdCustomerFindCustomerInfoRequest() *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest {
	return &AlibabaScbpAdCustomerFindCustomerInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.customer.find.customer.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopContextDto is TopContextDto Setter
// 用户信息
func (r *AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) SetTopContextDto(_topContextDto *TopContextDto) error {
	r._topContextDto = _topContextDto
	r.Set("top_context_dto", _topContextDto)
	return nil
}

// GetTopContextDto TopContextDto Getter
func (r AlibabaScbpAdCustomerFindCustomerInfoAPIRequest) GetTopContextDto() *TopContextDto {
	return r._topContextDto
}
