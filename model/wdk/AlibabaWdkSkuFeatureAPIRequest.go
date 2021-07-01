package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuFeatureAPIRequest
商品标记接口 API请求
alibaba.wdk.sku.feature

给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。 */
type AlibabaWdkSkuFeatureAPIRequest struct {
	model.Params
	// SkuFeatureDo
	_param *SkuFeatureDo
}

// New
