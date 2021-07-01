package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallFansCashpoolCreateAPIRequest
创建资金池 API请求
tmall.fans.cashpool.create

商家创建资金池接口 */
type TmallFansCashpoolCreateAPIRequest struct {
	model.Params
	// 创建资奖池输入对象
	_createCashPoolParamDo *CreateCashPoolParamDo
}

// New
