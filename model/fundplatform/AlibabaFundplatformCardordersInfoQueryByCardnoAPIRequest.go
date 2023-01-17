package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest 通过卡号查询卡信息 API请求
// alibaba.fundplatform.cardorders.info.query.by.cardno
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
type AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest struct {
	model.Params
	// 请求结构体
	_parameters *CardMakingInfoQueryRequest
}

// NewAlibabaFundplatformCardordersInfoQueryByCardnoRequest 初始化AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest对象
func NewAlibabaFundplatformCardordersInfoQueryByCardnoRequest() *AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest {
	return &AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.info.query.by.cardno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameters is Parameters Setter
// 请求结构体
func (r *AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest) SetParameters(_parameters *CardMakingInfoQueryRequest) error {
	r._parameters = _parameters
	r.Set("parameters", _parameters)
	return nil
}

// GetParameters Parameters Getter
func (r AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest) GetParameters() *CardMakingInfoQueryRequest {
	return r._parameters
}
