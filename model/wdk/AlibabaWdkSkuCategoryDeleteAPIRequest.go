package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryDeleteAPIRequest 商家类目删除接口 API请求
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
type AlibabaWdkSkuCategoryDeleteAPIRequest struct {
	model.Params
	// 类目删除请求模型
	_param *CategoryDo
}

// NewAlibabaWdkSkuCategoryDeleteRequest 初始化AlibabaWdkSkuCategoryDeleteAPIRequest对象
func NewAlibabaWdkSkuCategoryDeleteRequest() *AlibabaWdkSkuCategoryDeleteAPIRequest {
	return &AlibabaWdkSkuCategoryDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuCategoryDeleteAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 类目删除请求模型
func (r *AlibabaWdkSkuCategoryDeleteAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetParam() *CategoryDo {
	return r._param
}

var poolAlibabaWdkSkuCategoryDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuCategoryDeleteRequest()
	},
}

// GetAlibabaWdkSkuCategoryDeleteRequest 从 sync.Pool 获取 AlibabaWdkSkuCategoryDeleteAPIRequest
func GetAlibabaWdkSkuCategoryDeleteAPIRequest() *AlibabaWdkSkuCategoryDeleteAPIRequest {
	return poolAlibabaWdkSkuCategoryDeleteAPIRequest.Get().(*AlibabaWdkSkuCategoryDeleteAPIRequest)
}

// ReleaseAlibabaWdkSkuCategoryDeleteAPIRequest 将 AlibabaWdkSkuCategoryDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuCategoryDeleteAPIRequest(v *AlibabaWdkSkuCategoryDeleteAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuCategoryDeleteAPIRequest.Put(v)
}
