package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelProductsQueryAPIRequest
渠道中心-查询产品列表 API请求
tmall.channel.products.query

渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息 */
type TmallChannelProductsQueryAPIRequest struct {
	model.Params
	// 商家产品编码
	_productNumber string
	// 商家SKU编码
	_skuNumber string
	// 产品Id
	_productIds []int64
	// 分页大小
	_pageSize int64
	// 产品线Id
	_productLineId int64
	// 页码数，从1开始
	_pageNum int64
}

// New
