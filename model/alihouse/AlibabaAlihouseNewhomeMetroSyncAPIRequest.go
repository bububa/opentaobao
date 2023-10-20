package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomemetrosyncAPIRequest 地铁数据同步 API请求
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
type AlibabaalihousenewhomemetrosyncAPIRequest struct {
	model.Params
	// 地铁入参数据
	_baseMetroLineDto *BaseMetroLineDto
}

// NewAlibabaalihousenewhomemetrosyncRequest 初始化AlibabaalihousenewhomemetrosyncAPIRequest对象
func NewAlibabaalihousenewhomemetrosyncRequest() *AlibabaalihousenewhomemetrosyncAPIRequest {
	return &AlibabaalihousenewhomemetrosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomemetrosyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.metro.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomemetrosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomemetrosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseMetroLineDto is BaseMetroLineDto Setter
// 地铁入参数据
func (r *AlibabaalihousenewhomemetrosyncAPIRequest) SetBaseMetroLineDto(_baseMetroLineDto *BaseMetroLineDto) error {
	r._baseMetroLineDto = _baseMetroLineDto
	r.Set("base_metro_line_dto", _baseMetroLineDto)
	return nil
}

// GetBaseMetroLineDto BaseMetroLineDto Getter
func (r AlibabaalihousenewhomemetrosyncAPIRequest) GetBaseMetroLineDto() *BaseMetroLineDto {
	return r._baseMetroLineDto
}
