package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomebusinesssyncAPIRequest 商圈数据同步 API请求
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
type AlibabaalihousenewhomebusinesssyncAPIRequest struct {
	model.Params
	// 入参数据
	_baseBusinessDto *BaseBusinessDto
}

// NewAlibabaalihousenewhomebusinesssyncRequest 初始化AlibabaalihousenewhomebusinesssyncAPIRequest对象
func NewAlibabaalihousenewhomebusinesssyncRequest() *AlibabaalihousenewhomebusinesssyncAPIRequest {
	return &AlibabaalihousenewhomebusinesssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomebusinesssyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.business.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomebusinesssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomebusinesssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseBusinessDto is BaseBusinessDto Setter
// 入参数据
func (r *AlibabaalihousenewhomebusinesssyncAPIRequest) SetBaseBusinessDto(_baseBusinessDto *BaseBusinessDto) error {
	r._baseBusinessDto = _baseBusinessDto
	r.Set("base_business_dto", _baseBusinessDto)
	return nil
}

// GetBaseBusinessDto BaseBusinessDto Getter
func (r AlibabaalihousenewhomebusinesssyncAPIRequest) GetBaseBusinessDto() *BaseBusinessDto {
	return r._baseBusinessDto
}
