package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusSendgoodsCancelAPIRequest
取消发货 API请求
taobao.rdc.aligenius.sendgoods.cancel

提供商家在仅退款中发送取消发货状态 */
type TaobaoRdcAligeniusSendgoodsCancelAPIRequest struct {
	model.Params
	// 请求参数
	_param *CancelGoodsDto
}

// New
