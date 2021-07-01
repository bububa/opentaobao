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

// NewAlibabaWdkSkuFeatureRequest 初始化AlibabaWdkSkuFeatureAPIRequest对象
func NewAlibabaWdkSkuFeatureRequest() *AlibabaWdkSkuFeatureAPIRequest {
	return &AlibabaWdkSkuFeatureAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuFeatureAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.feature"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuFeatureAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// SkuFeatureDo
func (r *AlibabaWdkSkuFeatureAPIRequest) SetParam(_param *SkuFeatureDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkSkuFeatureAPIRequest) GetParam() *SkuFeatureDo {
	return r._param
}
