package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkShopQueryAPIRequest
门店查询接口 API请求
alibaba.wdk.shop.query

根据门店code查询门店信息 */
type AlibabaWdkShopQueryAPIRequest struct {
	model.Params
	// 如果不传，返回所有
	_ouCode string
}

// NewAlibabaWdkShopQueryRequest 初始化AlibabaWdkShopQueryAPIRequest对象
func NewAlibabaWdkShopQueryRequest() *AlibabaWdkShopQueryAPIRequest {
	return &AlibabaWdkShopQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkShopQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.shop.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkShopQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuCode Setter
// 如果不传，返回所有
func (r *AlibabaWdkShopQueryAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// Get OuCode Getter
func (r AlibabaWdkShopQueryAPIRequest) GetOuCode() string {
	return r._ouCode
}
