package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseOrderPerformAPIRequest 闲鱼验货担保服务商订单履约V1 API请求
// alibaba.idle.appraise.order.perform
//
// 闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
type AlibabaIdleAppraiseOrderPerformAPIRequest struct {
	model.Params
	// AppraiseOrderSynDto
	_appraiseOrderSynDto *AppraiseOrderSynDto
}

// NewAlibabaIdleAppraiseOrderPerformRequest 初始化AlibabaIdleAppraiseOrderPerformAPIRequest对象
func NewAlibabaIdleAppraiseOrderPerformRequest() *AlibabaIdleAppraiseOrderPerformAPIRequest {
	return &AlibabaIdleAppraiseOrderPerformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppraiseOrderSynDto Setter
// AppraiseOrderSynDto
func (r *AlibabaIdleAppraiseOrderPerformAPIRequest) SetAppraiseOrderSynDto(_appraiseOrderSynDto *AppraiseOrderSynDto) error {
	r._appraiseOrderSynDto = _appraiseOrderSynDto
	r.Set("appraise_order_syn_dto", _appraiseOrderSynDto)
	return nil
}

// Get AppraiseOrderSynDto Getter
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetAppraiseOrderSynDto() *AppraiseOrderSynDto {
	return r._appraiseOrderSynDto
}
