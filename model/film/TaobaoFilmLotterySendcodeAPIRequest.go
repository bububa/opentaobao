package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmLotterySendcodeAPIRequest
淘票票外部直发券 API请求
taobao.film.lottery.sendcode

淘票票外部直发券 */
type TaobaoFilmLotterySendcodeAPIRequest struct {
	model.Params
	// 外部商户发码请求
	_paramFCodeMerchantSendCodeRequest *FCodeMerchantSendCodeRq
}

// New
