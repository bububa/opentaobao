package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStorecollectConsumeAPIRequest
全渠道门店自提核销订单 API请求
taobao.omniorder.storecollect.consume

全渠道门店自提核销订单 */
type TaobaoOmniorderStorecollectConsumeAPIRequest struct {
	model.Params
	// 核销码
	_code string
	// 淘宝主订单ID
	_mainOrderId int64
	// 核销操作人信息
	_operator string
}

// New
