package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomwttopentradecreateorderAPIRequest 充值送活动下单接口 API请求
// alibaba.alicom.wtt.opentrade.createorder
//
// 提供给话费宝创建淘宝订单
type AlibabaalicomwttopentradecreateorderAPIRequest struct {
	model.Params
	// 入参请求说明
	_param0 *OpentradCreateOrderRequestDto
}

// NewAlibabaalicomwttopentradecreateorderRequest 初始化AlibabaalicomwttopentradecreateorderAPIRequest对象
func NewAlibabaalicomwttopentradecreateorderRequest() *AlibabaalicomwttopentradecreateorderAPIRequest {
	return &AlibabaalicomwttopentradecreateorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomwttopentradecreateorderAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.wtt.opentrade.createorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomwttopentradecreateorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomwttopentradecreateorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参请求说明
func (r *AlibabaalicomwttopentradecreateorderAPIRequest) SetParam0(_param0 *OpentradCreateOrderRequestDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaalicomwttopentradecreateorderAPIRequest) GetParam0() *OpentradCreateOrderRequestDto {
	return r._param0
}
