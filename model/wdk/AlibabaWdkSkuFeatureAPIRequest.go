package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuFeatureAPIRequest 商品标记接口 API请求
// alibaba.wdk.sku.feature
//
// 给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
type AlibabaWdkSkuFeatureAPIRequest struct {
	model.Params
	// SkuFeatureDo
	_param *SkuFeatureDo
}

// NewAlibabaWdkSkuFeatureRequest 初始化AlibabaWdkSkuFeatureAPIRequest对象
func NewAlibabaWdkSkuFeatureRequest() *AlibabaWdkSkuFeatureAPIRequest {
	return &AlibabaWdkSkuFeatureAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuFeatureAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuFeatureAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.feature"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuFeatureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuFeatureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// SkuFeatureDo
func (r *AlibabaWdkSkuFeatureAPIRequest) SetParam(_param *SkuFeatureDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuFeatureAPIRequest) GetParam() *SkuFeatureDo {
	return r._param
}

var poolAlibabaWdkSkuFeatureAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuFeatureRequest()
	},
}

// GetAlibabaWdkSkuFeatureRequest 从 sync.Pool 获取 AlibabaWdkSkuFeatureAPIRequest
func GetAlibabaWdkSkuFeatureAPIRequest() *AlibabaWdkSkuFeatureAPIRequest {
	return poolAlibabaWdkSkuFeatureAPIRequest.Get().(*AlibabaWdkSkuFeatureAPIRequest)
}

// ReleaseAlibabaWdkSkuFeatureAPIRequest 将 AlibabaWdkSkuFeatureAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuFeatureAPIRequest(v *AlibabaWdkSkuFeatureAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuFeatureAPIRequest.Put(v)
}
