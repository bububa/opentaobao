package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransPayRegisterAuditAPIRequest
阿信支付入驻审核通知 API请求
taobao.alitrip.axin.trans.pay.register.audit

阿信支付入驻审核通知 */
type TaobaoAlitripAxinTransPayRegisterAuditAPIRequest struct {
	model.Params
	// 支付入驻审核对象
	_axinPayRegisterAuditDto *AxinPayRegisterAuditDto
}

// New
