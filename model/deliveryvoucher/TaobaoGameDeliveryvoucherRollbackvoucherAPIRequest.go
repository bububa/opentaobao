package deliveryvoucher

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.rollbackvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) SetParam0(_param0 *RollbackVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest) GetParam0() *RollbackVoucherRequest {
	return r._param0
}
