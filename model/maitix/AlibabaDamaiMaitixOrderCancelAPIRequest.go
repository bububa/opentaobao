package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderCancelAPIRequest 大麦-库存释放 API请求
// alibaba.damai.maitix.order.cancel
//
// 库存释放
type AlibabaDamaiMaitixOrderCancelAPIRequest struct {
	model.Params
	// 库存释放入参
	_param *MoaUnlockTicketParam
}

// NewAlibabaDamaiMaitixOrderCancelRequest 初始化AlibabaDamaiMaitixOrderCancelAPIRequest对象
func NewAlibabaDamaiMaitixOrderCancelRequest() *AlibabaDamaiMaitixOrderCancelAPIRequest {
	return &AlibabaDamaiMaitixOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixOrderCancelAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 库存释放入参
func (r *AlibabaDamaiMaitixOrderCancelAPIRequest) SetParam(_param *MoaUnlockTicketParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetParam() *MoaUnlockTicketParam {
	return r._param
}

var poolAlibabaDamaiMaitixOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixOrderCancelRequest()
	},
}

// GetAlibabaDamaiMaitixOrderCancelRequest 从 sync.Pool 获取 AlibabaDamaiMaitixOrderCancelAPIRequest
func GetAlibabaDamaiMaitixOrderCancelAPIRequest() *AlibabaDamaiMaitixOrderCancelAPIRequest {
	return poolAlibabaDamaiMaitixOrderCancelAPIRequest.Get().(*AlibabaDamaiMaitixOrderCancelAPIRequest)
}

// ReleaseAlibabaDamaiMaitixOrderCancelAPIRequest 将 AlibabaDamaiMaitixOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderCancelAPIRequest(v *AlibabaDamaiMaitixOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderCancelAPIRequest.Put(v)
}
