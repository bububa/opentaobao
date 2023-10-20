package deliveryvoucher

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherSendvoucherAPIRequest 提货券发券接口 API请求
// taobao.game.deliveryvoucher.sendvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherSendvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// NewTaobaoGameDeliveryvoucherSendvoucherRequest 初始化TaobaoGameDeliveryvoucherSendvoucherAPIRequest对象
func NewTaobaoGameDeliveryvoucherSendvoucherRequest() *TaobaoGameDeliveryvoucherSendvoucherAPIRequest {
	return &TaobaoGameDeliveryvoucherSendvoucherAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGameDeliveryvoucherSendvoucherAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.sendvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherSendvoucherAPIRequest) SetParam0(_param0 *SendVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetParam0() *SendVoucherRequest {
	return r._param0
}

var poolTaobaoGameDeliveryvoucherSendvoucherAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGameDeliveryvoucherSendvoucherRequest()
	},
}

// GetTaobaoGameDeliveryvoucherSendvoucherRequest 从 sync.Pool 获取 TaobaoGameDeliveryvoucherSendvoucherAPIRequest
func GetTaobaoGameDeliveryvoucherSendvoucherAPIRequest() *TaobaoGameDeliveryvoucherSendvoucherAPIRequest {
	return poolTaobaoGameDeliveryvoucherSendvoucherAPIRequest.Get().(*TaobaoGameDeliveryvoucherSendvoucherAPIRequest)
}

// ReleaseTaobaoGameDeliveryvoucherSendvoucherAPIRequest 将 TaobaoGameDeliveryvoucherSendvoucherAPIRequest 放入 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherSendvoucherAPIRequest(v *TaobaoGameDeliveryvoucherSendvoucherAPIRequest) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherSendvoucherAPIRequest.Put(v)
}
