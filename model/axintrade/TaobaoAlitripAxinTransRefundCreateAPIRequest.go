package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransRefundCreateAPIRequest
阿信创建退款单 API请求
taobao.alitrip.axin.trans.refund.create

阿信供销平台-创建退款单服务 */
type TaobaoAlitripAxinTransRefundCreateAPIRequest struct {
	model.Params
	// 阿信创建退款单入参
	_axinRefundCreateDTO *AxinRefundCreateDto
}

// New
