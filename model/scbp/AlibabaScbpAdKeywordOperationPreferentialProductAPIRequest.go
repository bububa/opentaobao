package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordoperationpreferentialproductAPIRequest 操作优推品 API请求
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
type AlibabascbpadkeywordoperationpreferentialproductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 操作优推实体类
	_operationDto *CampaignTargetingWordSettingOperationDto
}

// NewAlibabascbpadkeywordoperationpreferentialproductRequest 初始化AlibabascbpadkeywordoperationpreferentialproductAPIRequest对象
func NewAlibabascbpadkeywordoperationpreferentialproductRequest() *AlibabascbpadkeywordoperationpreferentialproductAPIRequest {
	return &AlibabascbpadkeywordoperationpreferentialproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordoperationpreferentialproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.operation.preferential.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordoperationpreferentialproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordoperationpreferentialproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordoperationpreferentialproductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordoperationpreferentialproductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetOperationDto is OperationDto Setter
// 操作优推实体类
func (r *AlibabascbpadkeywordoperationpreferentialproductAPIRequest) SetOperationDto(_operationDto *CampaignTargetingWordSettingOperationDto) error {
	r._operationDto = _operationDto
	r.Set("operation_dto", _operationDto)
	return nil
}

// GetOperationDto OperationDto Getter
func (r AlibabascbpadkeywordoperationpreferentialproductAPIRequest) GetOperationDto() *CampaignTargetingWordSettingOperationDto {
	return r._operationDto
}
