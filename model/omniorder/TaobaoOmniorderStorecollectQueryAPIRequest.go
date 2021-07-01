package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStorecollectQueryAPIRequest
全渠道门店自提根据核销码查询订单 API请求
taobao.omniorder.storecollect.query

全渠道门店自提根据核销码查询订单 */
type TaobaoOmniorderStorecollectQueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// New
