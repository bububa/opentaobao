package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemCurrentpriceQueryAPIRequest
查询商品当前价格 API请求
alibaba.wdk.item.currentprice.query

通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量<=20,返回结果key为skuCode */
type AlibabaWdkItemCurrentpriceQueryAPIRequest struct {
	model.Params
	// 渠道店id
	_shopId int64
	// sku编码列表
	_skuCodes []string
	// 渠道
	_orderChannelCode string
}

// New
