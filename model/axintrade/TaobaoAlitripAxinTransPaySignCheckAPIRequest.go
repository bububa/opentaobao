package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransPaySignCheckAPIRequest
阿信支付宝验签服务 API请求
taobao.alitrip.axin.trans.pay.sign.check

阿信支付宝验签服务 */
type TaobaoAlitripAxinTransPaySignCheckAPIRequest struct {
	model.Params
	// 验签对象
	_axinPayCheckSignDto *AxinPayCheckSignDto
}

// New
