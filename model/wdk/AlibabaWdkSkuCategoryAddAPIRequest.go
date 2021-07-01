package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCategoryAddAPIRequest
商家类目新增接口 API请求
alibaba.wdk.sku.category.add

商家类目新增接口 */
type AlibabaWdkSkuCategoryAddAPIRequest struct {
	model.Params
	// 类目新增请求模型
	_param *CategoryDo
}

// New
