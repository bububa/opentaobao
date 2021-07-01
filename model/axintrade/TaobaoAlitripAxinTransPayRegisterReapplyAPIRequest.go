package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest
阿信支付入驻重新申请 API请求
taobao.alitrip.axin.trans.pay.register.reapply

阿信支付入驻重新申请
用于支付平台驳回，商户提交时的业务场景 */
type TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest struct {
	model.Params
	// 阿信支付入驻重新申请参数
	_axinPayRegisterCreateDTO *AxinPayRegisterCreateDto
}

// New
