package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantproducteditAPIRequest 商家产品服务-编辑产品 API请求
// alibaba.wdk.merchantproduct.edit
//
// 商家产品服务-编辑产品
type AlibabawdkmerchantproducteditAPIRequest struct {
	model.Params
	// 产品编辑入参
	_req *MerchantProductRequest
}

// NewAlibabawdkmerchantproducteditRequest 初始化AlibabawdkmerchantproducteditAPIRequest对象
func NewAlibabawdkmerchantproducteditRequest() *AlibabawdkmerchantproducteditAPIRequest {
	return &AlibabawdkmerchantproducteditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmerchantproducteditAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchantproduct.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmerchantproducteditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmerchantproducteditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 产品编辑入参
func (r *AlibabawdkmerchantproducteditAPIRequest) SetReq(_req *MerchantProductRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r AlibabawdkmerchantproducteditAPIRequest) GetReq() *MerchantProductRequest {
	return r._req
}
