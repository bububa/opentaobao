package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniDealerOdersListAPIRequest
全渠道经销商订单列表 API请求
taobao.omni.dealer.oders.list

全渠道经销商订单列表查询 */
type TaobaoOmniDealerOdersListAPIRequest struct {
	model.Params
	// 参数对象
	_queryParam *QueryOmniOrderRequest
}

// New
