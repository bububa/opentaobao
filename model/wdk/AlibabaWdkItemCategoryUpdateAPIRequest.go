package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemCategoryUpdateAPIRequest
修改类目 API请求
alibaba.wdk.item.category.update

修改类目 */
type AlibabaWdkItemCategoryUpdateAPIRequest struct {
	model.Params
	// 入参
	_bean *UpdateCategoryRequestBean
}

// New
