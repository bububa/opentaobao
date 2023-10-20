package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopCancelAPIRequest 阿里巴巴.天猫. 履约订单. 取消 API请求
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
type AlibabaTianmaoUopCancelAPIRequest struct {
	model.Params
	// 取消接口入参
	_hiErpCloseDto *HiErpCloseDto
}

// NewAlibabaTianmaoUopCancelRequest 初始化AlibabaTianmaoUopCancelAPIRequest对象
func NewAlibabaTianmaoUopCancelRequest() *AlibabaTianmaoUopCancelAPIRequest {
	return &AlibabaTianmaoUopCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTianmaoUopCancelAPIRequest) Reset() {
	r._hiErpCloseDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoUopCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoUopCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianmaoUopCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpCloseDto is HiErpCloseDto Setter
// 取消接口入参
func (r *AlibabaTianmaoUopCancelAPIRequest) SetHiErpCloseDto(_hiErpCloseDto *HiErpCloseDto) error {
	r._hiErpCloseDto = _hiErpCloseDto
	r.Set("hi_erp_close_dto", _hiErpCloseDto)
	return nil
}

// GetHiErpCloseDto HiErpCloseDto Getter
func (r AlibabaTianmaoUopCancelAPIRequest) GetHiErpCloseDto() *HiErpCloseDto {
	return r._hiErpCloseDto
}

var poolAlibabaTianmaoUopCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTianmaoUopCancelRequest()
	},
}

// GetAlibabaTianmaoUopCancelRequest 从 sync.Pool 获取 AlibabaTianmaoUopCancelAPIRequest
func GetAlibabaTianmaoUopCancelAPIRequest() *AlibabaTianmaoUopCancelAPIRequest {
	return poolAlibabaTianmaoUopCancelAPIRequest.Get().(*AlibabaTianmaoUopCancelAPIRequest)
}

// ReleaseAlibabaTianmaoUopCancelAPIRequest 将 AlibabaTianmaoUopCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaTianmaoUopCancelAPIRequest(v *AlibabaTianmaoUopCancelAPIRequest) {
	v.Reset()
	poolAlibabaTianmaoUopCancelAPIRequest.Put(v)
}
