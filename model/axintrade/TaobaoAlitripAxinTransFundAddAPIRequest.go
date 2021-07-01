package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundAddAPIRequest
创建资金单接口 API请求
taobao.alitrip.axin.trans.fund.add

创建资金单 */
type TaobaoAlitripAxinTransFundAddAPIRequest struct {
	model.Params
	// 创建资金单接口入参
	_axinFundCreateDTO *AxinFundCreateDto
}

// New
