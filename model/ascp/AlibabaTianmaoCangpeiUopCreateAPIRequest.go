package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoCangpeiUopCreateAPIRequest 阿里巴巴.天猫家装.仓配.履约订单.创建 API请求
// alibaba.tianmao.cangpei.uop.create
//
// 阿里巴巴.天猫家装.仓配.履约订单.创建
type AlibabaTianmaoCangpeiUopCreateAPIRequest struct {
	model.Params
	// ERP推单模型
	_hiErpOrderEvent *HiErpOrderEvent
}

// NewAlibabaTianmaoCangpeiUopCreateRequest 初始化AlibabaTianmaoCangpeiUopCreateAPIRequest对象
func NewAlibabaTianmaoCangpeiUopCreateRequest() *AlibabaTianmaoCangpeiUopCreateAPIRequest {
	return &AlibabaTianmaoCangpeiUopCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTianmaoCangpeiUopCreateAPIRequest) Reset() {
	r._hiErpOrderEvent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoCangpeiUopCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.cangpei.uop.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoCangpeiUopCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianmaoCangpeiUopCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpOrderEvent is HiErpOrderEvent Setter
// ERP推单模型
func (r *AlibabaTianmaoCangpeiUopCreateAPIRequest) SetHiErpOrderEvent(_hiErpOrderEvent *HiErpOrderEvent) error {
	r._hiErpOrderEvent = _hiErpOrderEvent
	r.Set("hi_erp_order_event", _hiErpOrderEvent)
	return nil
}

// GetHiErpOrderEvent HiErpOrderEvent Getter
func (r AlibabaTianmaoCangpeiUopCreateAPIRequest) GetHiErpOrderEvent() *HiErpOrderEvent {
	return r._hiErpOrderEvent
}

var poolAlibabaTianmaoCangpeiUopCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTianmaoCangpeiUopCreateRequest()
	},
}

// GetAlibabaTianmaoCangpeiUopCreateRequest 从 sync.Pool 获取 AlibabaTianmaoCangpeiUopCreateAPIRequest
func GetAlibabaTianmaoCangpeiUopCreateAPIRequest() *AlibabaTianmaoCangpeiUopCreateAPIRequest {
	return poolAlibabaTianmaoCangpeiUopCreateAPIRequest.Get().(*AlibabaTianmaoCangpeiUopCreateAPIRequest)
}

// ReleaseAlibabaTianmaoCangpeiUopCreateAPIRequest 将 AlibabaTianmaoCangpeiUopCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTianmaoCangpeiUopCreateAPIRequest(v *AlibabaTianmaoCangpeiUopCreateAPIRequest) {
	v.Reset()
	poolAlibabaTianmaoCangpeiUopCreateAPIRequest.Put(v)
}
