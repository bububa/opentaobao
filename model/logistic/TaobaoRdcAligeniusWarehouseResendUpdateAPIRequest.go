package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest
补发单状态回传 API请求
taobao.rdc.aligenius.warehouse.resend.update

补发单状态回传接口 */
type TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *UpdateResendStatusDto
}

// New
