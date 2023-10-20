package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaocangpeiuopcreateAPIRequest 阿里巴巴.天猫家装.仓配.履约订单.创建 API请求
// alibaba.tianmao.cangpei.uop.create
//
// 阿里巴巴.天猫家装.仓配.履约订单.创建
type AlibabatianmaocangpeiuopcreateAPIRequest struct {
	model.Params
	// ERP推单模型
	_hiErpOrderEvent *HiErpOrderEvent
}

// NewAlibabatianmaocangpeiuopcreateRequest 初始化AlibabatianmaocangpeiuopcreateAPIRequest对象
func NewAlibabatianmaocangpeiuopcreateRequest() *AlibabatianmaocangpeiuopcreateAPIRequest {
	return &AlibabatianmaocangpeiuopcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaocangpeiuopcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.cangpei.uop.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaocangpeiuopcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaocangpeiuopcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpOrderEvent is HiErpOrderEvent Setter
// ERP推单模型
func (r *AlibabatianmaocangpeiuopcreateAPIRequest) SetHiErpOrderEvent(_hiErpOrderEvent *HiErpOrderEvent) error {
	r._hiErpOrderEvent = _hiErpOrderEvent
	r.Set("hi_erp_order_event", _hiErpOrderEvent)
	return nil
}

// GetHiErpOrderEvent HiErpOrderEvent Getter
func (r AlibabatianmaocangpeiuopcreateAPIRequest) GetHiErpOrderEvent() *HiErpOrderEvent {
	return r._hiErpOrderEvent
}
