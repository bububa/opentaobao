package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuUpdateAPIRequest
更新商品 API请求
alibaba.wdk.sku.update

开放商品更新接口 */
type AlibabaWdkSkuUpdateAPIRequest struct {
	model.Params
	// 参数列表
	_paramList []SkuDo
}

// New
