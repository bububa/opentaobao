package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdConsumeAPIRequest 门店自送对码进行核销 API请求
// taobao.omniorder.dtd.consume
//
// 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoOmniorderDtdConsumeAPIRequest struct {
	model.Params
	// 核销信息
	_paramDoor2doorConsumeRequest *Door2doorConsumeRequest
}

// NewTaobaoOmniorderDtdConsumeRequest 初始化TaobaoOmniorderDtdConsumeAPIRequest对象
func NewTaobaoOmniorderDtdConsumeRequest() *TaobaoOmniorderDtdConsumeAPIRequest {
	return &TaobaoOmniorderDtdConsumeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderDtdConsumeAPIRequest) Reset() {
	r._paramDoor2doorConsumeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDoor2doorConsumeRequest is ParamDoor2doorConsumeRequest Setter
// 核销信息
func (r *TaobaoOmniorderDtdConsumeAPIRequest) SetParamDoor2doorConsumeRequest(_paramDoor2doorConsumeRequest *Door2doorConsumeRequest) error {
	r._paramDoor2doorConsumeRequest = _paramDoor2doorConsumeRequest
	r.Set("param_door2door_consume_request", _paramDoor2doorConsumeRequest)
	return nil
}

// GetParamDoor2doorConsumeRequest ParamDoor2doorConsumeRequest Getter
func (r TaobaoOmniorderDtdConsumeAPIRequest) GetParamDoor2doorConsumeRequest() *Door2doorConsumeRequest {
	return r._paramDoor2doorConsumeRequest
}

var poolTaobaoOmniorderDtdConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderDtdConsumeRequest()
	},
}

// GetTaobaoOmniorderDtdConsumeRequest 从 sync.Pool 获取 TaobaoOmniorderDtdConsumeAPIRequest
func GetTaobaoOmniorderDtdConsumeAPIRequest() *TaobaoOmniorderDtdConsumeAPIRequest {
	return poolTaobaoOmniorderDtdConsumeAPIRequest.Get().(*TaobaoOmniorderDtdConsumeAPIRequest)
}

// ReleaseTaobaoOmniorderDtdConsumeAPIRequest 将 TaobaoOmniorderDtdConsumeAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderDtdConsumeAPIRequest(v *TaobaoOmniorderDtdConsumeAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderDtdConsumeAPIRequest.Put(v)
}
