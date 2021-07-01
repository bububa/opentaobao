package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCombineskuQueryAPIRequest
组合商品查询接口 API请求
alibaba.wdk.sku.combinesku.query

查询组合商品接口 */
type AlibabaWdkSkuCombineskuQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param *SkuQueryDo
}

// New
