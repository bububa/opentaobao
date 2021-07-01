package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCategoryUpdateAPIRequest
商家类目修改接口 API请求
alibaba.wdk.sku.category.update

商家类目修改接口 */
type AlibabaWdkSkuCategoryUpdateAPIRequest struct {
	model.Params
	// 更新请求模型
	_param *CategoryDo
}

// New
