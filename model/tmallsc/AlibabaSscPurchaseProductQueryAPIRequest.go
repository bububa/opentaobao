package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscpurchaseproductqueryAPIRequest 查询已采购的服务产品 API请求
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
type AlibabasscpurchaseproductqueryAPIRequest struct {
	model.Params
}

// NewAlibabasscpurchaseproductqueryRequest 初始化AlibabasscpurchaseproductqueryAPIRequest对象
func NewAlibabasscpurchaseproductqueryRequest() *AlibabasscpurchaseproductqueryAPIRequest {
	return &AlibabasscpurchaseproductqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscpurchaseproductqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.purchase.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscpurchaseproductqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscpurchaseproductqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
