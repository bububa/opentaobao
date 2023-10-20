package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskufeatureAPIRequest 商品标记接口 API请求
// alibaba.wdk.sku.feature
//
// 给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
type AlibabawdkskufeatureAPIRequest struct {
	model.Params
	// SkuFeatureDo
	_param *SkuFeatureDo
}

// NewAlibabawdkskufeatureRequest 初始化AlibabawdkskufeatureAPIRequest对象
func NewAlibabawdkskufeatureRequest() *AlibabawdkskufeatureAPIRequest {
	return &AlibabawdkskufeatureAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskufeatureAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.feature"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskufeatureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskufeatureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// SkuFeatureDo
func (r *AlibabawdkskufeatureAPIRequest) SetParam(_param *SkuFeatureDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskufeatureAPIRequest) GetParam() *SkuFeatureDo {
	return r._param
}
