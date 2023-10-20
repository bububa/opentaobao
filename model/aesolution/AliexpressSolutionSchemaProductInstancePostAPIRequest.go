package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionschemaproductinstancepostAPIRequest aliexpress.solution.schema.product.instance.post API请求
// aliexpress.solution.schema.product.instance.post
//
// Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
type AliexpresssolutionschemaproductinstancepostAPIRequest struct {
	model.Params
	// Product instance data. The relative parameters description in schema json String are same as "aliexpress.solution.product.post" .Please note: the shipping_template_id should be replaced with your own shipping template id, which could be obtained through  https://developers.aliexpress.com/en/doc.htm?docId=43456&docType=2
	_productInstanceRequest string
	// More information of the request.
	_developerFeatures string
}

// NewAliexpresssolutionschemaproductinstancepostRequest 初始化AliexpresssolutionschemaproductinstancepostAPIRequest对象
func NewAliexpresssolutionschemaproductinstancepostRequest() *AliexpresssolutionschemaproductinstancepostAPIRequest {
	return &AliexpresssolutionschemaproductinstancepostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionschemaproductinstancepostAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.schema.product.instance.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionschemaproductinstancepostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionschemaproductinstancepostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductInstanceRequest is ProductInstanceRequest Setter
// Product instance data. The relative parameters description in schema json String are same as &#34;aliexpress.solution.product.post&#34; .Please note: the shipping_template_id should be replaced with your own shipping template id, which could be obtained through  https://developers.aliexpress.com/en/doc.htm?docId=43456&amp;docType=2
func (r *AliexpresssolutionschemaproductinstancepostAPIRequest) SetProductInstanceRequest(_productInstanceRequest string) error {
	r._productInstanceRequest = _productInstanceRequest
	r.Set("product_instance_request", _productInstanceRequest)
	return nil
}

// GetProductInstanceRequest ProductInstanceRequest Getter
func (r AliexpresssolutionschemaproductinstancepostAPIRequest) GetProductInstanceRequest() string {
	return r._productInstanceRequest
}

// SetDeveloperFeatures is DeveloperFeatures Setter
// More information of the request.
func (r *AliexpresssolutionschemaproductinstancepostAPIRequest) SetDeveloperFeatures(_developerFeatures string) error {
	r._developerFeatures = _developerFeatures
	r.Set("developer_features", _developerFeatures)
	return nil
}

// GetDeveloperFeatures DeveloperFeatures Getter
func (r AliexpresssolutionschemaproductinstancepostAPIRequest) GetDeveloperFeatures() string {
	return r._developerFeatures
}
