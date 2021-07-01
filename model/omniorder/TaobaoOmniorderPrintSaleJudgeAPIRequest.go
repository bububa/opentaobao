package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderPrintSaleJudgeAPIRequest
导购员判断 API请求
taobao.omniorder.print.sale.judge

用于判断当前子账号是否导购员 */
type TaobaoOmniorderPrintSaleJudgeAPIRequest struct {
	model.Params
	// 用户子账号ID
	_subUid int64
}

// New
