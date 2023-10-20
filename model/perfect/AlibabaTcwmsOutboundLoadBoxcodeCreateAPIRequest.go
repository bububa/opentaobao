package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatcwmsoutboundloadboxcodecreateAPIRequest 创建箱号 API请求
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
type AlibabatcwmsoutboundloadboxcodecreateAPIRequest struct {
	model.Params
	// 入口参数
	_boxCodeRequest *BoxCodeRequest
}

// NewAlibabatcwmsoutboundloadboxcodecreateRequest 初始化AlibabatcwmsoutboundloadboxcodecreateAPIRequest对象
func NewAlibabatcwmsoutboundloadboxcodecreateRequest() *AlibabatcwmsoutboundloadboxcodecreateAPIRequest {
	return &AlibabatcwmsoutboundloadboxcodecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatcwmsoutboundloadboxcodecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.load.boxcode.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatcwmsoutboundloadboxcodecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatcwmsoutboundloadboxcodecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBoxCodeRequest is BoxCodeRequest Setter
// 入口参数
func (r *AlibabatcwmsoutboundloadboxcodecreateAPIRequest) SetBoxCodeRequest(_boxCodeRequest *BoxCodeRequest) error {
	r._boxCodeRequest = _boxCodeRequest
	r.Set("box_code_request", _boxCodeRequest)
	return nil
}

// GetBoxCodeRequest BoxCodeRequest Getter
func (r AlibabatcwmsoutboundloadboxcodecreateAPIRequest) GetBoxCodeRequest() *BoxCodeRequest {
	return r._boxCodeRequest
}
