package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdtdconsumeAPIRequest 门店自送对码进行核销 API请求
// taobao.omniorder.dtd.consume
//
// 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoomniorderdtdconsumeAPIRequest struct {
	model.Params
	// 核销信息
	_paramDoor2doorConsumeRequest *Door2doorConsumeRequest
}

// NewTaobaoomniorderdtdconsumeRequest 初始化TaobaoomniorderdtdconsumeAPIRequest对象
func NewTaobaoomniorderdtdconsumeRequest() *TaobaoomniorderdtdconsumeAPIRequest {
	return &TaobaoomniorderdtdconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderdtdconsumeAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderdtdconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderdtdconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDoor2doorConsumeRequest is ParamDoor2doorConsumeRequest Setter
// 核销信息
func (r *TaobaoomniorderdtdconsumeAPIRequest) SetParamDoor2doorConsumeRequest(_paramDoor2doorConsumeRequest *Door2doorConsumeRequest) error {
	r._paramDoor2doorConsumeRequest = _paramDoor2doorConsumeRequest
	r.Set("param_door2door_consume_request", _paramDoor2doorConsumeRequest)
	return nil
}

// GetParamDoor2doorConsumeRequest ParamDoor2doorConsumeRequest Getter
func (r TaobaoomniorderdtdconsumeAPIRequest) GetParamDoor2doorConsumeRequest() *Door2doorConsumeRequest {
	return r._paramDoor2doorConsumeRequest
}
