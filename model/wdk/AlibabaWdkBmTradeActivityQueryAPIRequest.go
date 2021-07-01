package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkBmTradeActivityQueryAPIRequest
品牌营销的订单活动信息查询 API请求
alibaba.wdk.bm.trade.activity.query

品牌营销的订单活动信息查询 */
type AlibabaWdkBmTradeActivityQueryAPIRequest struct {
	model.Params
	// 入参
	_queryParam *IsvOrderQueryParam
}

// New
