package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoLanpeiUopCreateAPIRequest 阿里巴巴.天猫家装.揽配.履约订单.创建 API请求
// alibaba.tianmao.lanpei.uop.create
//
// 阿里巴巴.天猫家装.揽配.履约订单.创建
type AlibabaTianmaoLanpeiUopCreateAPIRequest struct {
	model.Params
	// ERP推单模型
	_hiErpOrderEvent *HiErpOrderEvent
}

// NewAlibabaTianmaoLanpeiUopCreateRequest 初始化AlibabaTianmaoLanpeiUopCreateAPIRequest对象
func NewAlibabaTianmaoLanpeiUopCreateRequest() *AlibabaTianmaoLanpeiUopCreateAPIRequest {
	return &AlibabaTianmaoLanpeiUopCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTianmaoLanpeiUopCreateAPIRequest) Reset() {
	r._hiErpOrderEvent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoLanpeiUopCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.lanpei.uop.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoLanpeiUopCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianmaoLanpeiUopCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpOrderEvent is HiErpOrderEvent Setter
// ERP推单模型
func (r *AlibabaTianmaoLanpeiUopCreateAPIRequest) SetHiErpOrderEvent(_hiErpOrderEvent *HiErpOrderEvent) error {
	r._hiErpOrderEvent = _hiErpOrderEvent
	r.Set("hi_erp_order_event", _hiErpOrderEvent)
	return nil
}

// GetHiErpOrderEvent HiErpOrderEvent Getter
func (r AlibabaTianmaoLanpeiUopCreateAPIRequest) GetHiErpOrderEvent() *HiErpOrderEvent {
	return r._hiErpOrderEvent
}

var poolAlibabaTianmaoLanpeiUopCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTianmaoLanpeiUopCreateRequest()
	},
}

// GetAlibabaTianmaoLanpeiUopCreateRequest 从 sync.Pool 获取 AlibabaTianmaoLanpeiUopCreateAPIRequest
func GetAlibabaTianmaoLanpeiUopCreateAPIRequest() *AlibabaTianmaoLanpeiUopCreateAPIRequest {
	return poolAlibabaTianmaoLanpeiUopCreateAPIRequest.Get().(*AlibabaTianmaoLanpeiUopCreateAPIRequest)
}

// ReleaseAlibabaTianmaoLanpeiUopCreateAPIRequest 将 AlibabaTianmaoLanpeiUopCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTianmaoLanpeiUopCreateAPIRequest(v *AlibabaTianmaoLanpeiUopCreateAPIRequest) {
	v.Reset()
	poolAlibabaTianmaoLanpeiUopCreateAPIRequest.Put(v)
}
