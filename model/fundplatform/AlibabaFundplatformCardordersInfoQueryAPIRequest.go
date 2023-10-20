package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardordersInfoQueryAPIRequest 根据制卡单分页查询卡信息 API请求
// alibaba.fundplatform.cardorders.info.query
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
type AlibabaFundplatformCardordersInfoQueryAPIRequest struct {
	model.Params
	// 请求结构体
	_parameters *CardMakingInfoQueryRequest
}

// NewAlibabaFundplatformCardordersInfoQueryRequest 初始化AlibabaFundplatformCardordersInfoQueryAPIRequest对象
func NewAlibabaFundplatformCardordersInfoQueryRequest() *AlibabaFundplatformCardordersInfoQueryAPIRequest {
	return &AlibabaFundplatformCardordersInfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardordersInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求结构体
func (r *AlibabaFundplatformCardordersInfoQueryAPIRequest) SetParameters(_parameters *CardMakingInfoQueryRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabaFundplatformCardordersInfoQueryAPIRequest) GetParameters() *CardMakingInfoQueryRequest {
	return r._parameters
}
