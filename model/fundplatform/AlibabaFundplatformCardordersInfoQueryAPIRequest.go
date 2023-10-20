package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordersinfoqueryAPIRequest 根据制卡单分页查询卡信息 API请求
// alibaba.fundplatform.cardorders.info.query
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
type AlibabafundplatformcardordersinfoqueryAPIRequest struct {
	model.Params
	// 请求结构体
	_parameters *CardMakingInfoQueryRequest
}

// NewAlibabafundplatformcardordersinfoqueryRequest 初始化AlibabafundplatformcardordersinfoqueryAPIRequest对象
func NewAlibabafundplatformcardordersinfoqueryRequest() *AlibabafundplatformcardordersinfoqueryAPIRequest {
	return &AlibabafundplatformcardordersinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardordersinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardordersinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardordersinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求结构体
func (r *AlibabafundplatformcardordersinfoqueryAPIRequest) SetParameters(_parameters *CardMakingInfoQueryRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabafundplatformcardordersinfoqueryAPIRequest) GetParameters() *CardMakingInfoQueryRequest {
	return r._parameters
}
