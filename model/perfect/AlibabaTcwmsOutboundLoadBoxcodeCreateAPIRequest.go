package perfect

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest 创建箱号 API请求
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
type AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest struct {
	model.Params
	// 入口参数
	_boxCodeRequest *BoxCodeRequest
}

// NewAlibabaTcwmsOutboundLoadBoxcodeCreateRequest 初始化AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest对象
func NewAlibabaTcwmsOutboundLoadBoxcodeCreateRequest() *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest {
	return &AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) Reset() {
	r._boxCodeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.load.boxcode.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBoxCodeRequest is BoxCodeRequest Setter
// 入口参数
func (r *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) SetBoxCodeRequest(_boxCodeRequest *BoxCodeRequest) error {
	r._boxCodeRequest = _boxCodeRequest
	r.Set("box_code_request", _boxCodeRequest)
	return nil
}

// GetBoxCodeRequest BoxCodeRequest Getter
func (r AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) GetBoxCodeRequest() *BoxCodeRequest {
	return r._boxCodeRequest
}

var poolAlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTcwmsOutboundLoadBoxcodeCreateRequest()
	},
}

// GetAlibabaTcwmsOutboundLoadBoxcodeCreateRequest 从 sync.Pool 获取 AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest
func GetAlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest() *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest {
	return poolAlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest.Get().(*AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest)
}

// ReleaseAlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest 将 AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest(v *AlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest) {
	v.Reset()
	poolAlibabaTcwmsOutboundLoadBoxcodeCreateAPIRequest.Put(v)
}
