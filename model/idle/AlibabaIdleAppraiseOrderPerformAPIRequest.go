package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleappraiseorderperformAPIRequest 闲鱼验货宝服务商订单履约 API请求
// alibaba.idle.appraise.order.perform
//
// 闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
type AlibabaidleappraiseorderperformAPIRequest struct {
	model.Params
	// AppraiseOrderSynDto
	_appraiseOrderSynDto *AppraiseOrderSynDto
}

// NewAlibabaidleappraiseorderperformRequest 初始化AlibabaidleappraiseorderperformAPIRequest对象
func NewAlibabaidleappraiseorderperformRequest() *AlibabaidleappraiseorderperformAPIRequest {
	return &AlibabaidleappraiseorderperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleappraiseorderperformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleappraiseorderperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleappraiseorderperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppraiseOrderSynDto is AppraiseOrderSynDto Setter
// AppraiseOrderSynDto
func (r *AlibabaidleappraiseorderperformAPIRequest) SetAppraiseOrderSynDto(_appraiseOrderSynDto *AppraiseOrderSynDto) error {
	r._appraiseOrderSynDto = _appraiseOrderSynDto
	r.Set("appraise_order_syn_dto", _appraiseOrderSynDto)
	return nil
}

// GetAppraiseOrderSynDto AppraiseOrderSynDto Getter
func (r AlibabaidleappraiseorderperformAPIRequest) GetAppraiseOrderSynDto() *AppraiseOrderSynDto {
	return r._appraiseOrderSynDto
}
