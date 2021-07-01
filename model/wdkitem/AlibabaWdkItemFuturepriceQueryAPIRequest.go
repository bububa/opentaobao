package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemFuturepriceQueryAPIRequest
单个商品未来价查询接口 API请求
alibaba.wdk.item.futureprice.query

查询单个商品未来价，融合了未来基础售价+未来促销价 */
type AlibabaWdkItemFuturepriceQueryAPIRequest struct {
	model.Params
	// 渠道店id
	_shopId int64
	// 商品编码
	_skuCode string
	// 渠道
	_orderChannelCode string
	// 开始时间
	_startTime string
	// 结束时间，结束时间-开始时间不能超过48小时
	_endTime string
}

// New
