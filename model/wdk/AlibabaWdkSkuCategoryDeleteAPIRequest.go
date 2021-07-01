package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCategoryDeleteAPIRequest
商家类目删除接口 API请求
alibaba.wdk.sku.category.delete

商家类目删除接口 */
type AlibabaWdkSkuCategoryDeleteAPIRequest struct {
	model.Params
	// 类目删除请求模型
	_param *CategoryDo
}

// New
