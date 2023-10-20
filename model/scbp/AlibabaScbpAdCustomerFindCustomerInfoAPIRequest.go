package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcustomerfindcustomerinfoAPIRequest 查询客户信息 API请求
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
type AlibabascbpadcustomerfindcustomerinfoAPIRequest struct {
	model.Params
	// 用户信息
	_topContextDto *TopContextDto
}

// NewAlibabascbpadcustomerfindcustomerinfoRequest 初始化AlibabascbpadcustomerfindcustomerinfoAPIRequest对象
func NewAlibabascbpadcustomerfindcustomerinfoRequest() *AlibabascbpadcustomerfindcustomerinfoAPIRequest {
	return &AlibabascbpadcustomerfindcustomerinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcustomerfindcustomerinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.customer.find.customer.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcustomerfindcustomerinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcustomerfindcustomerinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContextDto is TopContextDto Setter
// 用户信息
func (r *AlibabascbpadcustomerfindcustomerinfoAPIRequest) SetTopContextDto(_topContextDto *TopContextDto) error {
	r._topContextDto = _topContextDto
	r.Set("top_context_dto", _topContextDto)
	return nil
}

// GetTopContextDto TopContextDto Getter
func (r AlibabascbpadcustomerfindcustomerinfoAPIRequest) GetTopContextDto() *TopContextDto {
	return r._topContextDto
}
