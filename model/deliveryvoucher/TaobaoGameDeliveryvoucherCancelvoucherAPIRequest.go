package deliveryvoucher

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherCancelvoucherAPIRequest 作废券 API请求
// taobao.game.deliveryvoucher.cancelvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherCancelvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *CancelVoucherRequest
}

// NewTaobaoGameDeliveryvoucherCancelvoucherRequest 初始化TaobaoGameDeliveryvoucherCancelvoucherAPIRequest对象
func NewTaobaoGameDeliveryvoucherCancelvoucherRequest() *TaobaoGameDeliveryvoucherCancelvoucherAPIRequest {
	return &TaobaoGameDeliveryvoucherCancelvoucherAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.cancelvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) SetParam0(_param0 *CancelVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) GetParam0() *CancelVoucherRequest {
	return r._param0
}

var poolTaobaoGameDeliveryvoucherCancelvoucherAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGameDeliveryvoucherCancelvoucherRequest()
	},
}

// GetTaobaoGameDeliveryvoucherCancelvoucherRequest 从 sync.Pool 获取 TaobaoGameDeliveryvoucherCancelvoucherAPIRequest
func GetTaobaoGameDeliveryvoucherCancelvoucherAPIRequest() *TaobaoGameDeliveryvoucherCancelvoucherAPIRequest {
	return poolTaobaoGameDeliveryvoucherCancelvoucherAPIRequest.Get().(*TaobaoGameDeliveryvoucherCancelvoucherAPIRequest)
}

// ReleaseTaobaoGameDeliveryvoucherCancelvoucherAPIRequest 将 TaobaoGameDeliveryvoucherCancelvoucherAPIRequest 放入 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherCancelvoucherAPIRequest(v *TaobaoGameDeliveryvoucherCancelvoucherAPIRequest) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherCancelvoucherAPIRequest.Put(v)
}
