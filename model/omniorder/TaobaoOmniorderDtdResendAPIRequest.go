package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdtdresendAPIRequest 门店自送重发码 API请求
// taobao.omniorder.dtd.resend
//
// 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoomniorderdtdresendAPIRequest struct {
	model.Params
	// 淘宝主订单ID
	_mainOrderId int64
}

// NewTaobaoomniorderdtdresendRequest 初始化TaobaoomniorderdtdresendAPIRequest对象
func NewTaobaoomniorderdtdresendRequest() *TaobaoomniorderdtdresendAPIRequest {
	return &TaobaoomniorderdtdresendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderdtdresendAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderdtdresendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderdtdresendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝主订单ID
func (r *TaobaoomniorderdtdresendAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoomniorderdtdresendAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
