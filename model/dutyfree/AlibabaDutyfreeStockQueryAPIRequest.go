package dutyfree

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadutyfreestockqueryAPIRequest 对外库存查询接口 API请求
// alibaba.dutyfree.stock.query
//
// 对外部服务提供库存查询接口
type AlibabadutyfreestockqueryAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
}

// NewAlibabadutyfreestockqueryRequest 初始化AlibabadutyfreestockqueryAPIRequest对象
func NewAlibabadutyfreestockqueryRequest() *AlibabadutyfreestockqueryAPIRequest {
	return &AlibabadutyfreestockqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadutyfreestockqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dutyfree.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadutyfreestockqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadutyfreestockqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBarCode is BarCode Setter
// 条形码
func (r *AlibabadutyfreestockqueryAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r AlibabadutyfreestockqueryAPIRequest) GetBarCode() string {
	return r._barCode
}
