package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransPayRegisterCreateAPIRequest
提交支付服务开通 API请求
taobao.alitrip.axin.trans.pay.register.create

阿信供销平台-提交支付服务开通接口 */
type TaobaoAlitripAxinTransPayRegisterCreateAPIRequest struct {
	model.Params
	// 提交支付服务开通接口入参
	_createDTO *AxinPayRegisterCreateDto
}

// New
