package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantproductEditAPIRequest
商家产品服务-编辑产品 API请求
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品 */
type AlibabaWdkMerchantproductEditAPIRequest struct {
	model.Params
	// 产品编辑入参
	_req *MerchantProductRequest
}

// NewAlibabaWdkMerchantproductEditRequest 初始化AlibabaWdkMerchantproductEditAPIRequest对象
func NewAlibabaWdkMerchantproductEditRequest() *AlibabaWdkMerchantproductEditAPIRequest {
	return &AlibabaWdkMerchantproductEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantproductEditAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchantproduct.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantproductEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Req Setter
// 产品编辑入参
func (r *AlibabaWdkMerchantproductEditAPIRequest) SetReq(_req *MerchantProductRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// Get Req Getter
func (r AlibabaWdkMerchantproductEditAPIRequest) GetReq() *MerchantProductRequest {
	return r._req
}
