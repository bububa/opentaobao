package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaolanpeiuopcreateAPIRequest 阿里巴巴.天猫家装.揽配.履约订单.创建 API请求
// alibaba.tianmao.lanpei.uop.create
//
// 阿里巴巴.天猫家装.揽配.履约订单.创建
type AlibabatianmaolanpeiuopcreateAPIRequest struct {
	model.Params
	// ERP推单模型
	_hiErpOrderEvent *HiErpOrderEvent
}

// NewAlibabatianmaolanpeiuopcreateRequest 初始化AlibabatianmaolanpeiuopcreateAPIRequest对象
func NewAlibabatianmaolanpeiuopcreateRequest() *AlibabatianmaolanpeiuopcreateAPIRequest {
	return &AlibabatianmaolanpeiuopcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaolanpeiuopcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.lanpei.uop.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaolanpeiuopcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaolanpeiuopcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpOrderEvent is HiErpOrderEvent Setter
// ERP推单模型
func (r *AlibabatianmaolanpeiuopcreateAPIRequest) SetHiErpOrderEvent(_hiErpOrderEvent *HiErpOrderEvent) error {
	r._hiErpOrderEvent = _hiErpOrderEvent
	r.Set("hi_erp_order_event", _hiErpOrderEvent)
	return nil
}

// GetHiErpOrderEvent HiErpOrderEvent Getter
func (r AlibabatianmaolanpeiuopcreateAPIRequest) GetHiErpOrderEvent() *HiErpOrderEvent {
	return r._hiErpOrderEvent
}
