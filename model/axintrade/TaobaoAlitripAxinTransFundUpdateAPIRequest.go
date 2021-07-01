package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundUpdateAPIRequest
修改资金单接口 API请求
taobao.alitrip.axin.trans.fund.update

阿信供销平台-修改资金单接口 */
type TaobaoAlitripAxinTransFundUpdateAPIRequest struct {
	model.Params
	// 更新资金单接口入参
	_axinFundUpdateDTO *AxinFundUpdateDto
}

// New
