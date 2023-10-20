package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusRefundsCheckAPIRequest 退款信息审核 API请求
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
type TaobaoRdcAligeniusRefundsCheckAPIRequest struct {
	model.Params
	// 入参
	_param *RefundCheckDto
}

// NewTaobaoRdcAligeniusRefundsCheckRequest 初始化TaobaoRdcAligeniusRefundsCheckAPIRequest对象
func NewTaobaoRdcAligeniusRefundsCheckRequest() *TaobaoRdcAligeniusRefundsCheckAPIRequest {
	return &TaobaoRdcAligeniusRefundsCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusRefundsCheckAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusRefundsCheckAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.refunds.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusRefundsCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusRefundsCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoRdcAligeniusRefundsCheckAPIRequest) SetParam(_param *RefundCheckDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoRdcAligeniusRefundsCheckAPIRequest) GetParam() *RefundCheckDto {
	return r._param
}

var poolTaobaoRdcAligeniusRefundsCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusRefundsCheckRequest()
	},
}

// GetTaobaoRdcAligeniusRefundsCheckRequest 从 sync.Pool 获取 TaobaoRdcAligeniusRefundsCheckAPIRequest
func GetTaobaoRdcAligeniusRefundsCheckAPIRequest() *TaobaoRdcAligeniusRefundsCheckAPIRequest {
	return poolTaobaoRdcAligeniusRefundsCheckAPIRequest.Get().(*TaobaoRdcAligeniusRefundsCheckAPIRequest)
}

// ReleaseTaobaoRdcAligeniusRefundsCheckAPIRequest 将 TaobaoRdcAligeniusRefundsCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusRefundsCheckAPIRequest(v *TaobaoRdcAligeniusRefundsCheckAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusRefundsCheckAPIRequest.Put(v)
}
