package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpPurchasePriceCreateAPIRequest ascp采购价写入接口 API请求
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
type AlibabaAscpPurchasePriceCreateAPIRequest struct {
	model.Params
	// 采购价创建/更新请求
	_createRequest *AlibabaAscpPurchasePriceCreateRequest
}

// NewAlibabaAscpPurchasePriceCreateRequest 初始化AlibabaAscpPurchasePriceCreateAPIRequest对象
func NewAlibabaAscpPurchasePriceCreateRequest() *AlibabaAscpPurchasePriceCreateAPIRequest {
	return &AlibabaAscpPurchasePriceCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.purchase.price.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CreateRequest Setter
// 采购价创建/更新请求
func (r *AlibabaAscpPurchasePriceCreateAPIRequest) SetCreateRequest(_createRequest *AlibabaAscpPurchasePriceCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// Get CreateRequest Getter
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetCreateRequest() *AlibabaAscpPurchasePriceCreateRequest {
	return r._createRequest
}
