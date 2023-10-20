package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordersinfoquerybycardnoAPIRequest 通过卡号查询卡信息 API请求
// alibaba.fundplatform.cardorders.info.query.by.cardno
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
type AlibabafundplatformcardordersinfoquerybycardnoAPIRequest struct {
	model.Params
	// 请求结构体
	_parameters *CardMakingInfoQueryRequest
}

// NewAlibabafundplatformcardordersinfoquerybycardnoRequest 初始化AlibabafundplatformcardordersinfoquerybycardnoAPIRequest对象
func NewAlibabafundplatformcardordersinfoquerybycardnoRequest() *AlibabafundplatformcardordersinfoquerybycardnoAPIRequest {
	return &AlibabafundplatformcardordersinfoquerybycardnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardordersinfoquerybycardnoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.info.query.by.cardno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardordersinfoquerybycardnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardordersinfoquerybycardnoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求结构体
func (r *AlibabafundplatformcardordersinfoquerybycardnoAPIRequest) SetParameters(_parameters *CardMakingInfoQueryRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabafundplatformcardordersinfoquerybycardnoAPIRequest) GetParameters() *CardMakingInfoQueryRequest {
	return r._parameters
}
