package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCombineskuUpdateAPIRequest
组合商品更新接口 API请求
alibaba.wdk.sku.combinesku.update

组合商品修改接口 */
type AlibabaWdkSkuCombineskuUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// New
