package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryQueryAPIRequest 商家类目获取接口 API请求
// alibaba.wdk.sku.category.query
//
// 商家类目获取接口
type AlibabaWdkSkuCategoryQueryAPIRequest struct {
	model.Params
	// 查询类目请模型
	_param *CategoryDo
}

// NewAlibabaWdkSkuCategoryQueryRequest 初始化AlibabaWdkSkuCategoryQueryAPIRequest对象
func NewAlibabaWdkSkuCategoryQueryRequest() *AlibabaWdkSkuCategoryQueryAPIRequest {
	return &AlibabaWdkSkuCategoryQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuCategoryQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询类目请模型
func (r *AlibabaWdkSkuCategoryQueryAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetParam() *CategoryDo {
	return r._param
}

var poolAlibabaWdkSkuCategoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuCategoryQueryRequest()
	},
}

// GetAlibabaWdkSkuCategoryQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuCategoryQueryAPIRequest
func GetAlibabaWdkSkuCategoryQueryAPIRequest() *AlibabaWdkSkuCategoryQueryAPIRequest {
	return poolAlibabaWdkSkuCategoryQueryAPIRequest.Get().(*AlibabaWdkSkuCategoryQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuCategoryQueryAPIRequest 将 AlibabaWdkSkuCategoryQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuCategoryQueryAPIRequest(v *AlibabaWdkSkuCategoryQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuCategoryQueryAPIRequest.Put(v)
}
