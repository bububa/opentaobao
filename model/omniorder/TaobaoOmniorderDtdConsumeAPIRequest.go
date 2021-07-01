package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderDtdConsumeAPIRequest
门店自送对码进行核销 API请求
taobao.omniorder.dtd.consume

该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。 */
type TaobaoOmniorderDtdConsumeAPIRequest struct {
	model.Params
	// 核销信息
	_paramDoor2doorConsumeRequest *Door2doorConsumeRequest
}

// NewTaobaoOmniorderDtdConsumeRequest 初始化TaobaoOmniorderDtdConsumeAPIRequest对象
func NewTaobaoOmniorderDtdConsumeRequest() *TaobaoOmniorderDtdConsumeAPIRequest {
	return &TaobaoOmniorderDtdConsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamDoor2doorConsumeRequest Setter
// 核销信息
func (r *TaobaoOmniorderDtdConsumeAPIRequest) SetParamDoor2doorConsumeRequest(_paramDoor2doorConsumeRequest *Door2doorConsumeRequest) error {
	r._paramDoor2doorConsumeRequest = _paramDoor2doorConsumeRequest
	r.Set("param_door2door_consume_request", _paramDoor2doorConsumeRequest)
	return nil
}

// Get ParamDoor2doorConsumeRequest Getter
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetParamDoor2doorConsumeRequest() *Door2doorConsumeRequest {
	return r._paramDoor2doorConsumeRequest
}
