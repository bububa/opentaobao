package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryUpdateAPIRequest 商家类目修改接口 API请求
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
type AlibabaWdkSkuCategoryUpdateAPIRequest struct {
	model.Params
	// 更新请求模型
	_param *CategoryDo
}

// NewAlibabaWdkSkuCategoryUpdateRequest 初始化AlibabaWdkSkuCategoryUpdateAPIRequest对象
func NewAlibabaWdkSkuCategoryUpdateRequest() *AlibabaWdkSkuCategoryUpdateAPIRequest {
	return &AlibabaWdkSkuCategoryUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuCategoryUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuCategoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新请求模型
func (r *AlibabaWdkSkuCategoryUpdateAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuCategoryUpdateAPIRequest) GetParam() *CategoryDo {
	return r._param
}

var poolAlibabaWdkSkuCategoryUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuCategoryUpdateRequest()
	},
}

// GetAlibabaWdkSkuCategoryUpdateRequest 从 sync.Pool 获取 AlibabaWdkSkuCategoryUpdateAPIRequest
func GetAlibabaWdkSkuCategoryUpdateAPIRequest() *AlibabaWdkSkuCategoryUpdateAPIRequest {
	return poolAlibabaWdkSkuCategoryUpdateAPIRequest.Get().(*AlibabaWdkSkuCategoryUpdateAPIRequest)
}

// ReleaseAlibabaWdkSkuCategoryUpdateAPIRequest 将 AlibabaWdkSkuCategoryUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuCategoryUpdateAPIRequest(v *AlibabaWdkSkuCategoryUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuCategoryUpdateAPIRequest.Put(v)
}
