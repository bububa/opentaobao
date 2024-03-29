package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseOrderPerformAPIRequest 闲鱼验货宝服务商订单履约 API请求
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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleAppraiseOrderPerformAPIRequest) Reset() {
	r._appraiseOrderSynDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppraiseOrderSynDto is AppraiseOrderSynDto Setter
// AppraiseOrderSynDto
func (r *AlibabaIdleAppraiseOrderPerformAPIRequest) SetAppraiseOrderSynDto(_appraiseOrderSynDto *AppraiseOrderSynDto) error {
	r._appraiseOrderSynDto = _appraiseOrderSynDto
	r.Set("appraise_order_syn_dto", _appraiseOrderSynDto)
	return nil
}

// GetAppraiseOrderSynDto AppraiseOrderSynDto Getter
func (r AlibabaIdleAppraiseOrderPerformAPIRequest) GetAppraiseOrderSynDto() *AppraiseOrderSynDto {
	return r._appraiseOrderSynDto
}

var poolAlibabaIdleAppraiseOrderPerformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleAppraiseOrderPerformRequest()
	},
}

// GetAlibabaIdleAppraiseOrderPerformRequest 从 sync.Pool 获取 AlibabaIdleAppraiseOrderPerformAPIRequest
func GetAlibabaIdleAppraiseOrderPerformAPIRequest() *AlibabaIdleAppraiseOrderPerformAPIRequest {
	return poolAlibabaIdleAppraiseOrderPerformAPIRequest.Get().(*AlibabaIdleAppraiseOrderPerformAPIRequest)
}

// ReleaseAlibabaIdleAppraiseOrderPerformAPIRequest 将 AlibabaIdleAppraiseOrderPerformAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleAppraiseOrderPerformAPIRequest(v *AlibabaIdleAppraiseOrderPerformAPIRequest) {
	v.Reset()
	poolAlibabaIdleAppraiseOrderPerformAPIRequest.Put(v)
}
