package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuAddAPIRequest 新增商品 API请求
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
type AlibabaWdkSkuAddAPIRequest struct {
	model.Params
	// 商品列表
	_paramList []SkuDo
}

// NewAlibabaWdkSkuAddRequest 初始化AlibabaWdkSkuAddAPIRequest对象
func NewAlibabaWdkSkuAddRequest() *AlibabaWdkSkuAddAPIRequest {
	return &AlibabaWdkSkuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuAddAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 商品列表
func (r *AlibabaWdkSkuAddAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuAddAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}

var poolAlibabaWdkSkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuAddRequest()
	},
}

// GetAlibabaWdkSkuAddRequest 从 sync.Pool 获取 AlibabaWdkSkuAddAPIRequest
func GetAlibabaWdkSkuAddAPIRequest() *AlibabaWdkSkuAddAPIRequest {
	return poolAlibabaWdkSkuAddAPIRequest.Get().(*AlibabaWdkSkuAddAPIRequest)
}

// ReleaseAlibabaWdkSkuAddAPIRequest 将 AlibabaWdkSkuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuAddAPIRequest(v *AlibabaWdkSkuAddAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuAddAPIRequest.Put(v)
}
