package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPosTradeQueryAPIRequest
轻pos品牌营销查询接口 API请求
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息 */
type AlibabaWdkPosTradeQueryAPIRequest struct {
	model.Params
	// 查询请求
	_queryRequest *FastBuyPosQueryRequest
}

// New
