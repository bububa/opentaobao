package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuUpdateAPIRequest 更新商品 API请求
// alibaba.wdk.sku.update
//
// 开放商品更新接口
type AlibabaWdkSkuUpdateAPIRequest struct {
	model.Params
	// 参数列表
	_paramList []SkuDo
}

// NewAlibabaWdkSkuUpdateRequest 初始化AlibabaWdkSkuUpdateAPIRequest对象
func NewAlibabaWdkSkuUpdateRequest() *AlibabaWdkSkuUpdateAPIRequest {
	return &AlibabaWdkSkuUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuUpdateAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 参数列表
func (r *AlibabaWdkSkuUpdateAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuUpdateAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}

var poolAlibabaWdkSkuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuUpdateRequest()
	},
}

// GetAlibabaWdkSkuUpdateRequest 从 sync.Pool 获取 AlibabaWdkSkuUpdateAPIRequest
func GetAlibabaWdkSkuUpdateAPIRequest() *AlibabaWdkSkuUpdateAPIRequest {
	return poolAlibabaWdkSkuUpdateAPIRequest.Get().(*AlibabaWdkSkuUpdateAPIRequest)
}

// ReleaseAlibabaWdkSkuUpdateAPIRequest 将 AlibabaWdkSkuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuUpdateAPIRequest(v *AlibabaWdkSkuUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuUpdateAPIRequest.Put(v)
}
