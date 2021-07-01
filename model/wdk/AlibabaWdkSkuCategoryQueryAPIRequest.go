package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCategoryQueryAPIRequest
商家类目获取接口 API请求
alibaba.wdk.sku.category.query

商家类目获取接口 */
type AlibabaWdkSkuCategoryQueryAPIRequest struct {
	model.Params
	// 查询类目请模型
	_param *CategoryDo
}

// New
