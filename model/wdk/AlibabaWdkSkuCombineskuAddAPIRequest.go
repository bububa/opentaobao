package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCombineskuAddAPIRequest
组合商品新增接口 API请求
alibaba.wdk.sku.combinesku.add

组合商品新增接口 */
type AlibabaWdkSkuCombineskuAddAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// New
