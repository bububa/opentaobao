package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantproductEditAPIRequest 商家产品服务-编辑产品 API请求
// alibaba.wdk.merchantproduct.edit
//
// 商家产品服务-编辑产品
type AlibabaWdkMerchantproductEditAPIRequest struct {
	model.Params
	// 产品编辑入参
	_req *MerchantProductRequest
}

// NewAlibabaWdkMerchantproductEditRequest 初始化AlibabaWdkMerchantproductEditAPIRequest对象
func NewAlibabaWdkMerchantproductEditRequest() *AlibabaWdkMerchantproductEditAPIRequest {
	return &AlibabaWdkMerchantproductEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMerchantproductEditAPIRequest) Reset() {
	r._req = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantproductEditAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchantproduct.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantproductEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantproductEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 产品编辑入参
func (r *AlibabaWdkMerchantproductEditAPIRequest) SetReq(_req *MerchantProductRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r AlibabaWdkMerchantproductEditAPIRequest) GetReq() *MerchantProductRequest {
	return r._req
}

var poolAlibabaWdkMerchantproductEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMerchantproductEditRequest()
	},
}

// GetAlibabaWdkMerchantproductEditRequest 从 sync.Pool 获取 AlibabaWdkMerchantproductEditAPIRequest
func GetAlibabaWdkMerchantproductEditAPIRequest() *AlibabaWdkMerchantproductEditAPIRequest {
	return poolAlibabaWdkMerchantproductEditAPIRequest.Get().(*AlibabaWdkMerchantproductEditAPIRequest)
}

// ReleaseAlibabaWdkMerchantproductEditAPIRequest 将 AlibabaWdkMerchantproductEditAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMerchantproductEditAPIRequest(v *AlibabaWdkMerchantproductEditAPIRequest) {
	v.Reset()
	poolAlibabaWdkMerchantproductEditAPIRequest.Put(v)
}
