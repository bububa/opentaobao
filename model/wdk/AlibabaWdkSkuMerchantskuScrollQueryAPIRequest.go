package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuMerchantskuScrollQueryAPIRequest
商家商品批量查询接口 API请求
alibaba.wdk.sku.merchantsku.scroll.query

提供主档商品数据接口查询 */
type AlibabaWdkSkuMerchantskuScrollQueryAPIRequest struct {
	model.Params
	// HM
	_orgNo string
	// 第一次为null，后面从结果中获取
	_scrollId string
}

// New
