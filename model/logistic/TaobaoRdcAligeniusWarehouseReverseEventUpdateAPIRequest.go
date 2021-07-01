package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest
销退单事件回传接口 API请求
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台 */
type TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *ReverseEventInfoDto
}

// New
