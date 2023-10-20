package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvoucherrollbackvoucherAPIRequest 回滚券 API请求
// taobao.game.deliveryvoucher.rollbackvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaogamedeliveryvoucherrollbackvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *RollbackVoucherRequest
}

// NewTaobaogamedeliveryvoucherrollbackvoucherRequest 初始化TaobaogamedeliveryvoucherrollbackvoucherAPIRequest对象
func NewTaobaogamedeliveryvoucherrollbackvoucherRequest() *TaobaogamedeliveryvoucherrollbackvoucherAPIRequest {
	return &TaobaogamedeliveryvoucherrollbackvoucherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogamedeliveryvoucherrollbackvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.rollbackvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogamedeliveryvoucherrollbackvoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogamedeliveryvoucherrollbackvoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaogamedeliveryvoucherrollbackvoucherAPIRequest) SetParam0(_param0 *RollbackVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaogamedeliveryvoucherrollbackvoucherAPIRequest) GetParam0() *RollbackVoucherRequest {
	return r._param0
}
