package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundQueryByOrderAPIRequest
通过外部订单ID查询所有资金单 API请求
taobao.alitrip.axin.trans.fund.query.by.order

阿信供销平台-通过外部订单ID查询所有资金单 */
type TaobaoAlitripAxinTransFundQueryByOrderAPIRequest struct {
	model.Params
	// 入参
	_paramAxinFundListQueryDTO *AxinFundListQueryDto
}

// New
