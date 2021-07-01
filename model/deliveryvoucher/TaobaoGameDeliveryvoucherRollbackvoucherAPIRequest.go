package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest
回滚券 API请求
taobao.game.deliveryvoucher.rollbackvoucher

提货券发券接口：同步券和订单的关联信息 */
type TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *RollbackVoucherRequest
}

// New
