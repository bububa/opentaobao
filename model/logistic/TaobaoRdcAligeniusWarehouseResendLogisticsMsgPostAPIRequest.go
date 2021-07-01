package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest
补发单物流信息回传 API请求
taobao.rdc.aligenius.warehouse.resend.logistics.msg.post

补发单erp物流信息回传平台 */
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest struct {
	model.Params
	// 参数
	_param0 *SendResendLogisticsMsgDto
}

// New
