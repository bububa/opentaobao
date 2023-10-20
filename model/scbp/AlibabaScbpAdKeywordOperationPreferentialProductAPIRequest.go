package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) Reset() {
	r._topContext = nil
	r._operationDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.operation.preferential.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpAdKeywordOperationPreferentialProductAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordOperationPreferentialProductRequest()
	},
}

// GetAlibabaScbpAdKeywordOperationPreferentialProductRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest
func GetAlibabaScbpAdKeywordOperationPreferentialProductAPIRequest() *AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest {
	return poolAlibabaScbpAdKeywordOperationPreferentialProductAPIRequest.Get().(*AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordOperationPreferentialProductAPIRequest 将 AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordOperationPreferentialProductAPIRequest(v *AlibabaScbpAdKeywordOperationPreferentialProductAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordOperationPreferentialProductAPIRequest.Put(v)
}
