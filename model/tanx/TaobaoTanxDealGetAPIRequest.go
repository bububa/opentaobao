package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxDealGetAPIRequest
对外部dsp提供交易id查询接口 API请求
taobao.tanx.deal.get

对外部dsp提供交易id查询接口 */
type TaobaoTanxDealGetAPIRequest struct {
	model.Params
	// dsp用户id
	_dspId int64
	// 交易id
	_dealId int64
	// 1970年到现在的时间，毫秒
	_signTime int64
	// 验证token
	_token string
}

// New
