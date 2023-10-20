package deliveryvoucher

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherSendgoodsAPIRequest 提货券发货接口 API请求
// taobao.game.deliveryvoucher.sendgoods
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherSendgoodsAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// NewTaobaoGameDeliveryvoucherSendgoodsRequest 初始化TaobaoGameDeliveryvoucherSendgoodsAPIRequest对象
func NewTaobaoGameDeliveryvoucherSendgoodsRequest() *TaobaoGameDeliveryvoucherSendgoodsAPIRequest {
	return &TaobaoGameDeliveryvoucherSendgoodsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGameDeliveryvoucherSendgoodsAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.sendgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherSendgoodsAPIRequest) SetParam0(_param0 *SendVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetParam0() *SendVoucherRequest {
	return r._param0
}

var poolTaobaoGameDeliveryvoucherSendgoodsAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGameDeliveryvoucherSendgoodsRequest()
	},
}

// GetTaobaoGameDeliveryvoucherSendgoodsRequest 从 sync.Pool 获取 TaobaoGameDeliveryvoucherSendgoodsAPIRequest
func GetTaobaoGameDeliveryvoucherSendgoodsAPIRequest() *TaobaoGameDeliveryvoucherSendgoodsAPIRequest {
	return poolTaobaoGameDeliveryvoucherSendgoodsAPIRequest.Get().(*TaobaoGameDeliveryvoucherSendgoodsAPIRequest)
}

// ReleaseTaobaoGameDeliveryvoucherSendgoodsAPIRequest 将 TaobaoGameDeliveryvoucherSendgoodsAPIRequest 放入 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherSendgoodsAPIRequest(v *TaobaoGameDeliveryvoucherSendgoodsAPIRequest) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherSendgoodsAPIRequest.Put(v)
}
