package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingPriceAPIRequest
促销价签服务 API请求
alibaba.wdk.marketing.price

获取营销-促销商品中的实时价格 */
type AlibabaWdkMarketingPriceAPIRequest struct {
	model.Params
	// 单页大小
	_pageSize int64
	// 页码
	_pageIndex int64
	// 商品sku
	_skuCodes []string
	// 门店标识数组
	_shopIds []int64
	// 查询结束时间(sku_codes非空无效)
	_endTime string
	// 查询开始时间(sku_codes非空无效)
	_beginTime string
}

// New
