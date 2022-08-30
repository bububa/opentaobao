package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest 操作优推品 API请求
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
type AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 操作优推实体类
	_operationDto *CampaignTargetingWordSettingOperationDto
}

// NewAlibabaScbpAdKeywordOperationPreferentialProductRequest 初始化AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest对象
func NewAlibabaScbpAdKeywordOperationPreferentialProductRequest() *AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest {
	return &AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.operation.preferential.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetOperationDto is OperationDto Setter
// 操作优推实体类
func (r *AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) SetOperationDto(_operationDto *CampaignTargetingWordSettingOperationDto) error {
	r._operationDto = _operationDto
	r.Set("operation_dto", _operationDto)
	return nil
}

// GetOperationDto OperationDto Getter
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetOperationDto() *CampaignTargetingWordSettingOperationDto {
	return r._operationDto
}
