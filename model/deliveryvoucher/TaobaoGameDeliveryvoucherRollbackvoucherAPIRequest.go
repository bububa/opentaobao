package deliveryvoucher

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest 回滚券 API请求
// taobao.game.deliveryvoucher.rollbackvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *RollbackVoucherRequest
}

// NewTaobaoGameDeliveryvoucherRollbackvoucherRequest 初始化TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest对象
func NewTaobaoGameDeliveryvoucherRollbackvoucherRequest() *TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest {
	return &TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.rollbackvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) SetParam0(_param0 *RollbackVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetParam0() *RollbackVoucherRequest {
	return r._param0
}

var poolTaobaoGameDeliveryvoucherRollbackvoucherAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGameDeliveryvoucherRollbackvoucherRequest()
	},
}

// GetTaobaoGameDeliveryvoucherRollbackvoucherRequest 从 sync.Pool 获取 TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest
func GetTaobaoGameDeliveryvoucherRollbackvoucherAPIRequest() *TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest {
	return poolTaobaoGameDeliveryvoucherRollbackvoucherAPIRequest.Get().(*TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest)
}

// ReleaseTaobaoGameDeliveryvoucherRollbackvoucherAPIRequest 将 TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest 放入 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherRollbackvoucherAPIRequest(v *TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherRollbackvoucherAPIRequest.Put(v)
}
