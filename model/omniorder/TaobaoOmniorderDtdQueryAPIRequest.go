package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderDtdQueryAPIRequest
门店自送根据核销码查订单 API请求
taobao.omniorder.dtd.query

门店自送根据核销码码查询订单信息 */
type TaobaoOmniorderDtdQueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// New
