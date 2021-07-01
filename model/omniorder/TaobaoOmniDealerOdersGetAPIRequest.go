package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniDealerOdersGetAPIRequest
获取单笔全渠道经销商订单的详细信息 API请求
taobao.omni.dealer.oders.get

全渠道经销商获取单笔订单的详细信息 */
type TaobaoOmniDealerOdersGetAPIRequest struct {
	model.Params
	// 主订单ID
	_mainOrderId int64
}

// New
