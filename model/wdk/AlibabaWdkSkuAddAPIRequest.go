package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuAddAPIRequest
新增商品 API请求
alibaba.wdk.sku.add

创建RT门店商品或DC商品 */
type AlibabaWdkSkuAddAPIRequest struct {
	model.Params
	// 商品列表
	_paramList []SkuDo
}

// New
