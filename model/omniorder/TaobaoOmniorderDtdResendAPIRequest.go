package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdResendAPIRequest 门店自送重发码 API请求
// taobao.omniorder.dtd.resend
//
// 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoOmniorderDtdResendAPIRequest struct {
	model.Params
	// 淘宝主订单ID
	_mainOrderId int64
}

// NewTaobaoOmniorderDtdResendRequest 初始化TaobaoOmniorderDtdResendAPIRequest对象
func NewTaobaoOmniorderDtdResendRequest() *TaobaoOmniorderDtdResendAPIRequest {
	return &TaobaoOmniorderDtdResendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderDtdResendAPIRequest) Reset() {
	r._mainOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdResendAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdResendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderDtdResendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝主订单ID
func (r *TaobaoOmniorderDtdResendAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoOmniorderDtdResendAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

var poolTaobaoOmniorderDtdResendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderDtdResendRequest()
	},
}

// GetTaobaoOmniorderDtdResendRequest 从 sync.Pool 获取 TaobaoOmniorderDtdResendAPIRequest
func GetTaobaoOmniorderDtdResendAPIRequest() *TaobaoOmniorderDtdResendAPIRequest {
	return poolTaobaoOmniorderDtdResendAPIRequest.Get().(*TaobaoOmniorderDtdResendAPIRequest)
}

// ReleaseTaobaoOmniorderDtdResendAPIRequest 将 TaobaoOmniorderDtdResendAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderDtdResendAPIRequest(v *TaobaoOmniorderDtdResendAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderDtdResendAPIRequest.Put(v)
}
