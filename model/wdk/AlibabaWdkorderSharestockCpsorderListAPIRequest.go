package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockCpsorderListAPIRequest cps正向分销订单批量回流 API请求
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
type AlibabaWdkorderSharestockCpsorderListAPIRequest struct {
	model.Params
	// 入参
	_cpsOrderRequest *CpsOrderRequest
}

// NewAlibabaWdkorderSharestockCpsorderListRequest 初始化AlibabaWdkorderSharestockCpsorderListAPIRequest对象
func NewAlibabaWdkorderSharestockCpsorderListRequest() *AlibabaWdkorderSharestockCpsorderListAPIRequest {
	return &AlibabaWdkorderSharestockCpsorderListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.cpsorder.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpsOrderRequest is CpsOrderRequest Setter
// 入参
func (r *AlibabaWdkorderSharestockCpsorderListAPIRequest) SetCpsOrderRequest(_cpsOrderRequest *CpsOrderRequest) error {
	r._cpsOrderRequest = _cpsOrderRequest
	r.Set("cps_order_request", _cpsOrderRequest)
	return nil
}

// GetCpsOrderRequest CpsOrderRequest Getter
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetCpsOrderRequest() *CpsOrderRequest {
	return r._cpsOrderRequest
}
