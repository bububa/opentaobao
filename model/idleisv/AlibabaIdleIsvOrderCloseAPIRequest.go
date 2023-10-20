package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderCloseAPIRequest 服务商闲鱼卖家主动关闭订单 API请求
// alibaba.idle.isv.order.close
//
// 供外部服务商 isv 提供卖家主动关闭交易订单的功能
type AlibabaIdleIsvOrderCloseAPIRequest struct {
	model.Params
	// 输入参数
	_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto
}

// NewAlibabaIdleIsvOrderCloseRequest 初始化AlibabaIdleIsvOrderCloseAPIRequest对象
func NewAlibabaIdleIsvOrderCloseRequest() *AlibabaIdleIsvOrderCloseAPIRequest {
	return &AlibabaIdleIsvOrderCloseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvOrderCloseAPIRequest) Reset() {
	r._isvAppraiseIsvOrderCloseDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvAppraiseIsvOrderCloseDto is IsvAppraiseIsvOrderCloseDto Setter
// 输入参数
func (r *AlibabaIdleIsvOrderCloseAPIRequest) SetIsvAppraiseIsvOrderCloseDto(_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto) error {
	r._isvAppraiseIsvOrderCloseDto = _isvAppraiseIsvOrderCloseDto
	r.Set("isv_appraise_isv_order_close_dto", _isvAppraiseIsvOrderCloseDto)
	return nil
}

// GetIsvAppraiseIsvOrderCloseDto IsvAppraiseIsvOrderCloseDto Getter
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetIsvAppraiseIsvOrderCloseDto() *AppraiseIsvOrderCloseDto {
	return r._isvAppraiseIsvOrderCloseDto
}

var poolAlibabaIdleIsvOrderCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvOrderCloseRequest()
	},
}

// GetAlibabaIdleIsvOrderCloseRequest 从 sync.Pool 获取 AlibabaIdleIsvOrderCloseAPIRequest
func GetAlibabaIdleIsvOrderCloseAPIRequest() *AlibabaIdleIsvOrderCloseAPIRequest {
	return poolAlibabaIdleIsvOrderCloseAPIRequest.Get().(*AlibabaIdleIsvOrderCloseAPIRequest)
}

// ReleaseAlibabaIdleIsvOrderCloseAPIRequest 将 AlibabaIdleIsvOrderCloseAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvOrderCloseAPIRequest(v *AlibabaIdleIsvOrderCloseAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvOrderCloseAPIRequest.Put(v)
}
